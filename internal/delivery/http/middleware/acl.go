package middleware

import (
	"net/netip"

	"github.com/Citrus0974/ProxyProject/internal/models"
	"github.com/Citrus0974/ProxyProject/internal/usecase/acl"
	"github.com/gin-gonic/gin"
)

type ACLMiddleware struct {
	Service *acl.Service
}

func (a *ACLMiddleware) ACL() gin.HandlerFunc {

	return func(c *gin.Context) {

		ip, err := netip.ParseAddr(
			c.ClientIP(),
		)

		if err != nil {
			c.AbortWithStatus(403)
			return
		}

		decision := a.Service.Check(ip)

		if decision == models.Deny {
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}
