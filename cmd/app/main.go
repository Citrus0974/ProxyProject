package main

import (
	"fmt"
	"log"

	"github.com/Citrus0974/ProxyProject/config"
	"github.com/Citrus0974/ProxyProject/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app := cfg.App
	proxy := cfg.Proxy
	server := cfg.Server

	fmt.Println("\nStarting " + app.Name + " version " + app.Version)

	engine := gin.Default()
	controller.NewRouter(engine, proxy.BaseURL)
	engine.Run(":" + server.Port)
}
