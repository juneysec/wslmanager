package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"wslmanager/domains/domainobjects"
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

	c.JSON(http.StatusOK, adaptDistributions((workspace.Distributions)))
}

func (p *DistributionsAPI) DistributionsDistributionGet(c *gin.Context) {
	var uriParam uriParameter
	c.ShouldBindUri(&uriParam)

	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		log.Fatal(err)
	}

	dist, err := workspace.Find(uriParam.Distribution)
	if err != nil {
		c.JSON(http.StatusNotFound, schema.ResponseError{
			Code: "404",
			Message: "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, adaptDistribution(dist))
}

func (p *DistributionsAPI) DistributionsDistributionPost(c *gin.Context) {
	c.JSON(http.StatusNotFound, schema.ResponseError{
		Code:    "1001",
		Message: "サポートされていません。",
	})
}

func (p *DistributionsAPI) DistributionsDistributionPut(c *gin.Context) {
	var requestBody schema.RequestDistributionPut

	c.ShouldBind(&requestBody)
	distributionName := c.Param("distribution")

	if len(distributionName) == 0 {
		c.JSON(http.StatusBadRequest, schema.ResponseError{
			Code: "101",
			Message: "不正なリクエスト",
		})
		return
	}

	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	distribution, err := workspace.Find(distributionName)
	if err != nil {
		c.JSON(http.StatusNotFound, schema.ResponseError{
			Code:    "404",
			Message: err.Error(),
		})
		return
	}

	switch (requestBody.Command) {
	case "start":  // 起動
		err = workspace.Start(distribution.Name)
	case "stop":   // 停止
		err = workspace.Stop(distribution.Name)
	case "shell":  // シェル起動
		err = workspace.ExecShell(distribution.Name)
	case "export": // エクスポート
		err = workspace.Export(distribution.Name, requestBody.Path)
	case "set-default": // デフォルトに変更する
		err = workspace.SetDefault(distribution.Name)
	default:
		c.JSON(http.StatusBadRequest, schema.ResponseError{
			Code:    "400",
			Message: fmt.Sprintf("不正なコマンド %s", requestBody.Command),
		})
		return
	}

	if err != nil {
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

	c.JSON(http.StatusOK, adaptDistributions((workspace.Distributions)))
}

func (p *DistributionsAPI) DistributionsDistributionDelete(c *gin.Context) {
	c.JSON(http.StatusNotFound, schema.ResponseError{
		Code:    "1001",
		Message: "サポートされていません。",
	})
}

func adaptDistribution(dist *domainobjects.Distribution) *schema.ResponseDistribution {
	return &schema.ResponseDistribution{
		IsDefault: dist.IsDefault,
		Name:      dist.Name,
		State:     dist.State,
		Version:   dist.Version,
	}
}

func adaptDistributions(dist []*domainobjects.Distribution) []*schema.ResponseDistribution {
	var distributions []*schema.ResponseDistribution

	for _, d := range dist {
		distributions = append(distributions, adaptDistribution(d))
	}

	return distributions
}

