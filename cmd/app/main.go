package main

import (
	"fmt"
	"log"

	"github.com/Citrus0974/ProxyProject/config"
	"github.com/Citrus0974/ProxyProject/internal/delivery/http/handler"
	deliveryRouter "github.com/Citrus0974/ProxyProject/internal/delivery/http/router"
	proxyInfrastructure "github.com/Citrus0974/ProxyProject/internal/infrastructure/proxy"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatalf("Config error: %s \n", err)
	}

	app := cfg.App
	proxy := cfg.Proxy
	server := cfg.Server

	fmt.Printf("Starting %s version %s\n", app.Name, app.Version)

	reverseProxy, err := proxyInfrastructure.NewReverseProxy(proxy.BaseURL)
	if err != nil {
		log.Fatalf("Proxy err: %s \n", err)
	}

	proxyHandler := handler.NewProxyHandler(reverseProxy)
	healthHandler := handler.NewHealthHandler()

	engine := gin.New()

	deliveryRouter.RouterSetup(engine, deliveryRouter.Handlers{
		ProxyHandler:  proxyHandler,
		HealthHandler: healthHandler,
	})

	err = engine.Run(":" + server.Port)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
