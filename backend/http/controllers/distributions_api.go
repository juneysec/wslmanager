package controllers

import (
	"fmt"
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
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adaptDistributions((workspace.Distributions)))
}

func (p *DistributionsAPI) DistributionsDistributionGet(c *gin.Context) {
	var uriParam uriParameter
	c.ShouldBindUri(&uriParam)

	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	dist, err := workspace.Find(uriParam.Distribution)
	if err != nil {
		c.JSON(http.StatusNotFound, schema.ResponseError{
			Code:    "404",
			Message: "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, adaptDistribution(dist))
}

// ディストリビューションのインポート
func (p *DistributionsAPI) DistributionsPost(c *gin.Context) {
	var requestBody schema.RequestDistributionPost
	c.ShouldBind(&requestBody)

	workspace, err := workspaces.NewWSLManagerWorkspace()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ResponseError{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	if err := workspace.Import(requestBody.Name, requestBody.VhdPath, requestBody.SourcePath); err != nil {
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

func (p *DistributionsAPI) DistributionsDistributionPut(c *gin.Context) {
	var requestBody schema.RequestDistributionPut

	c.ShouldBind(&requestBody)
	distributionName := c.Param("distribution")

	if len(distributionName) == 0 {
		c.JSON(http.StatusBadRequest, schema.ResponseError{
			Code:    "101",
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

	switch requestBody.Command {
	case "start": // 起動
		err = workspace.Start(distribution.Name)
	case "stop": // 停止
		err = workspace.Stop(distribution.Name)
	case "shell": // シェル起動
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
	var requestBody schema.RequestDistributionDelete
	c.ShouldBind(&requestBody)
	distributionName := c.Param("distribution")

	if distributionName == requestBody.Name {
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

		if err := workspace.Unregister(distribution.Name); err != nil {
			c.JSON(http.StatusInternalServerError, schema.ResponseError{
				Code:    "500",
				Message: err.Error(),
			})
		}

		if err := workspace.Fetch(); err != nil {
			c.JSON(http.StatusInternalServerError, schema.ResponseError{
				Code:    "500",
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, adaptDistributions((workspace.Distributions)))
	} else {
		c.JSON(http.StatusNotFound, schema.ResponseError{
			Code:    "400",
			Message: "不正な削除処理を行おうとしました",
		})
	}
}

func adaptDistribution(dist *domainobjects.Distribution) *schema.ResponseDistribution {
	return &schema.ResponseDistribution{
		IsDefault: dist.IsDefault,
		Name:      dist.Name,
		State:     dist.State,
		Version:   dist.Version,
		VhdPath:   dist.VhdPath,
	}
}

func adaptDistributions(dist []*domainobjects.Distribution) []*schema.ResponseDistribution {
	var distributions []*schema.ResponseDistribution

	for _, d := range dist {
		distributions = append(distributions, adaptDistribution(d))
	}

	return distributions
}
