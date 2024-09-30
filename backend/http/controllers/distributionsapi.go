package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"wslmanager/domains/workspaces"
	"wslmanager/http/schema"
)

type DistributionsAPI struct {
}

type uriParameter struct {
	Distribution string `uri:"distribution" binding:"required"`
}

func (p *DistributionsAPI) DistributionsGet(c *gin.Context) {
	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		log.Fatal(err)
	}

	workspace.Fetch()
	var response []schema.ResponseDistribution
	for _, dist := range workspace.Distributions {
		response = append(response, schema.ResponseDistribution{
			IsDefault: dist.IsDefault,
			Name:      dist.Name,
			State:     dist.State,
			Version:   dist.Version,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (p *DistributionsAPI) DistributionsDistributionGet(c *gin.Context) {
	var uriParam uriParameter
	c.ShouldBindUri(&uriParam)

	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		log.Fatal(err)
	}

	workspace.Fetch()
	for _, dist := range workspace.Distributions {
		if dist.Name == uriParam.Distribution {
			c.JSON(http.StatusOK, schema.ResponseDistribution{
				IsDefault: dist.IsDefault,
				Name:      dist.Name,
				State:     dist.State,
				Version:   dist.Version,
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, schema.ResponseDistribution{})
}
