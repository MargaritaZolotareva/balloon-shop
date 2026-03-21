package middleware

import (
	"github.com/NYTimes/gziphandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GzipMiddleware() gin.HandlerFunc {
	handler := gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
		c.Next()
	}
}
