// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * WSLManager API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

package http

import (
  "github.com/gin-gonic/gin"
  "wslmanager/http/controllers"
)

type DistributionsAPIController interface { 
  DistributionsDistributionDelete(*gin.Context)
  DistributionsDistributionGet(*gin.Context)
  DistributionsDistributionPut(*gin.Context)
  DistributionsGet(*gin.Context)
  DistributionsPost(*gin.Context)
}

type OnlineDistributionsAPIController interface { 
  OnlineDistributionsGet(*gin.Context)
  OnlineDistributionsPost(*gin.Context)
}

func registerDistributionsAPIRoutes(e *gin.Engine, r DistributionsAPIController) {
  e.DELETE("/api/v1/distributions/:distribution", r.DistributionsDistributionDelete)
  e.GET("/api/v1/distributions/:distribution", r.DistributionsDistributionGet)
  e.PUT("/api/v1/distributions/:distribution", r.DistributionsDistributionPut)
  e.GET("/api/v1/distributions", r.DistributionsGet)
  e.POST("/api/v1/distributions", r.DistributionsPost)
}

func registerOnlineDistributionsAPIRoutes(e *gin.Engine, r OnlineDistributionsAPIController) {
  e.GET("/api/v1/online-distributions", r.OnlineDistributionsGet)
  e.POST("/api/v1/online-distributions", r.OnlineDistributionsPost)
}

func RegisterRoutes(e *gin.Engine) {
  registerDistributionsAPIRoutes(e, &controllers.DistributionsAPI{})
  registerOnlineDistributionsAPIRoutes(e, &controllers.OnlineDistributionsAPI{})
}
