package router

import (
	"github.com/Citrus0974/ProxyProject/internal/delivery/http/handler"
	"github.com/Citrus0974/ProxyProject/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ProxyHandler  *handler.ProxyHandler
	HealthHandler *handler.HealthHandler
}

type Middlewares struct {
	ACL middleware.ACLMiddleware
}

func RouterSetup(engine *gin.Engine, h Handlers) {
	//proxy managing api
	healthApi := engine.Group("/proxy")
	{
		healthApi.GET("/health", h.HealthHandler.Handle)
	}
	healthApi.Use(gin.Recovery())
	healthApi.Use(gin.Logger())

	//proxy
	engine.NoRoute(gin.Logger(), gin.Recovery(), h.ProxyHandler.Handle)
}
