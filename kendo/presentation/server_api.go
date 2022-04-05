package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/hollson/kendo/config"
	"github.com/hollson/kendo/presentation/api"
	"github.com/hollson/kendo/presentation/middleware"
)

var router = gin.Default()

func NewHttpServer() *gin.Engine {
	router.Use(middleware.Errors())
	// r.Use(middleware.Cors())

	if config.AppMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	// restfull: Put Post Get Delete
	f := router.Group("/file")
	{
		f.PUT("/", api.AddFileHandler)
		f.GET("/:id", api.GetFileHandler)
	}
	return router
}
