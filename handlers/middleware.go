package handlers

import (
	"github.com/gin-gonic/gin"
)

const myAPIKey = "eUbP9shywUygMx7u"

//AuthMiddleware use http header api key to verify the request for authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != myAPIKey {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}
