package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Citrus0974/ProxyProject/config"
	"github.com/Citrus0974/ProxyProject/internal/delivery/http/handler"
	deliveryRouter "github.com/Citrus0974/ProxyProject/internal/delivery/http/router"
	"github.com/Citrus0974/ProxyProject/internal/infrastructure/logger"
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

	logger := logger.NewZerologLogger()
	logger.Debugf("test message %d %s", 123, "ip")

	reverseProxy, err := proxyInfrastructure.NewReverseProxy(proxy.BaseURL)
	if err != nil {
		log.Fatalf("Proxy err: %s \n", err)
	}

	proxyHandler := handler.NewProxyHandler(reverseProxy)
	healthHandler := handler.NewHealthHandler()

	engine := gin.New()

	deliveryRouter.RouterSetup(engine, deliveryRouter.Handlers{
		HealthHandler: healthHandler,
		ProxyHandler:  proxyHandler,
	})

	err = engine.Run(":" + server.Port)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}

	go func() {
		config.UpdateConfig(cfg)
		time.Sleep(time.Duration(app.ReloadTimerSeconds) * time.Second)
	}()
}
