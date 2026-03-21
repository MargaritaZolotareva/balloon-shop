package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RedirectWWW() gin.HandlerFunc {
	return func(c *gin.Context) {
		domainName := os.Getenv("DOMAIN_NAME")
		host := c.Request.Host
		if host == ("www." + domainName) {
			target := "https://" + domainName + c.Request.RequestURI
			c.Redirect(301, target)
			c.Abort()
			return
		}
		c.Next()
	}
}
