package handler

import (
	"net/http/httputil"
	"github.com/gin-gonic/gin"
)

type ProxyHandler struct {
	proxy *httputil.ReverseProxy
}

func NewProxyHandler(proxy *httputil.ReverseProxy) *ProxyHandler {
	return &ProxyHandler{
		proxy: proxy,
	}
}

func (h *ProxyHandler) Handle(c *gin.Context) {
	h.proxy.ServeHTTP(c.Writer, c.Request)
}