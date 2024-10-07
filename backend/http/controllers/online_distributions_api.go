package controllers

import (
	"net/http"
	"wslmanager/domains/domainobjects"
	"wslmanager/domains/workspaces"
	"wslmanager/http/schema"

	"github.com/gin-gonic/gin"
)

type OnlineDistributionsAPI struct {
}

func (p *OnlineDistributionsAPI) OnlineDistributionsGet(c *gin.Context) {
	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	onlineDistributions, err := workspace.GetOnlineDistributions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adaptOnlineDistributions(onlineDistributions))
}

func (p *OnlineDistributionsAPI) OnlineDistributionsPost(c *gin.Context) {
	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	var requestBody schema.RequestOnlineDistributionPost
	c.ShouldBind(&requestBody)

	if err := workspace.InstallOnlineDistribution(requestBody.Name); err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	if err := workspace.Fetch(); err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adaptDistributions(workspace.Distributions))
}

func adaptOnlineDistributions(distributions []*domainobjects.OnlineDistribution) []*schema.ResponseOnlineDistribution {
	var retval []*schema.ResponseOnlineDistribution

	for _, distribution := range distributions {
		retval = append(retval, adaptOnlineDistribution(distribution))
	}
	
	return retval;
}

func adaptOnlineDistribution(d *domainobjects.OnlineDistribution) *schema.ResponseOnlineDistribution {
	return &schema.ResponseOnlineDistribution{
		Name: d.Name,
		FriendlyName: d.FriendlyName,
		AlreadyInstalled: d.AlreadyInstalled,
	}
}