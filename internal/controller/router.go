package controller

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewRouter(engine *gin.Engine, targetURL string) {
	api := engine.Group("/proxy")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	engine.NoRoute(func(c *gin.Context) {
		target, _ := url.Parse(targetURL)
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}
