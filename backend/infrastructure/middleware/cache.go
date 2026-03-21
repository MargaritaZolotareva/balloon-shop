package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	PhotoCacheTTL = 30 * 24 * time.Hour
	ShortCacheTTL = 30 * time.Second
)

type SimpleCacheItem struct {
	Content    []byte
	Expiration time.Time
}
type SafeCache struct {
	mu   sync.RWMutex
	Data map[string]*SimpleCacheItem
}

func NewSafeCache() *SafeCache {
	c := &SafeCache{
		Data: make(map[string]*SimpleCacheItem),
	}
	go c.cleanupLoop(5 * time.Minute)
	return c
}

func (c *SafeCache) cleanupLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		now := time.Now()
		c.mu.Lock()
		for k, v := range c.Data {
			if now.After(v.Expiration) {
				delete(c.Data, k)
			}
		}
		c.mu.Unlock()
	}
}

func SimpleCacheMiddleware(duration time.Duration, cache *SafeCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		key := c.Request.RequestURI

		cache.mu.RLock()
		item, ok := cache.Data[key]
		cache.mu.RUnlock()

		if ok && time.Now().Before(item.Expiration) {
			c.Writer.Header().Set("X-Cache", "HIT")
			c.Writer.Write(item.Content)
			c.Abort()
			return
		}

		writer := &bodyWriter{body: make([]byte, 0), ResponseWriter: c.Writer}
		c.Writer = writer
		c.Next()

		status := c.Writer.Status()
		body := writer.body

		if len(body) == 0 {
			return
		}

		var ttl time.Duration

		if strings.HasPrefix(key, "/photos/") {
			ttl = PhotoCacheTTL
		} else if status >= 500 {
			return
		} else if status == 404 || status == 400 {
			ttl = ShortCacheTTL
		} else {
			ttl = duration
		}

		cache.mu.Lock()
		cache.Data[key] = &SimpleCacheItem{
			Content:    writer.body,
			Expiration: time.Now().Add(ttl),
		}
		cache.mu.Unlock()
		c.Writer.Header().Set("X-Cache", "MISS")
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}
