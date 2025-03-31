package router

import (
	"time"
	"uala-challenge/config"
	_ "uala-challenge/docs"
	"uala-challenge/middleware"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(cfg *config.Config, getStatusHandler gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(requestid.New())
	router.Use(middleware.LoggerMiddleware())

	router.HandleMethodNotAllowed = true
	router.NoMethod(middleware.MethodNotAllowed)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/status", middleware.Timeout(time.Millisecond*time.Duration(cfg.GetStatusTimeout)), getStatusHandler)

	return router
}
