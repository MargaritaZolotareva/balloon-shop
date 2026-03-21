package middleware

import "github.com/gin-gonic/gin"

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'nonce-XIar3vBAdOctuDvWQ+2beyGikm8iaIoiCdVOjXwa9nY=' https://mc.yandex.ru;"+
			"style-src 'self' 'unsafe-inline'; img-src 'self' data:; connect-src 'self' https://api.telegram.org https://mc.yandex.ru wss://mc.yandex.ru https://mc.yandex.com;"+
			"frame-src https://www.google.com https://mc.yandex.ru https://yandex.ru/;")
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		c.Next()
	}
}
