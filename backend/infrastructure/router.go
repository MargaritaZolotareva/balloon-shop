package infrastructure

import (
	"backend/api"
	"backend/controllers"
	"backend/infrastructure/metrics"
	"backend/infrastructure/middleware"
	"backend/repositories"
	"backend/services"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

var (
	FrontendDistDir    = os.Getenv("FRONTEND_DIST_DIR")
	FrontendJSDir      = filepath.Join(FrontendDistDir, "js")
	FrontendCSSDir     = filepath.Join(FrontendDistDir, "css")
	PhotoCacheDuration = 7 * 24 * time.Hour
)

func SetupRouter(db *gorm.DB, botAPI *tgbotapi.BotAPI) *gin.Engine {
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService()
	productController := controllers.NewProductController(db, productRepo, productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryController := controllers.NewCategoryController(db, categoryRepo)

	homepageController := controllers.NewHomepageController(db, categoryRepo)

	messageService := services.NewMessageService(botAPI)
	messageController := controllers.NewMessageController(messageService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("VUE_APP_URL")},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
	r.Use(middleware.RedirectWWW())
	r.Use(middleware.GzipMiddleware())
	r.Use(middleware.SecurityHeadersMiddleware())

	apiPref := r.Group("/api")
	apiPref.GET("/metrics", gin.WrapH(promhttp.Handler()))
	apiPref.POST("/lead-form", messageController.SendLeadMessage)
	apiPref.GET("/products/:slug", productController.GetProductBySlug)
	apiPref.GET("/categories", categoryController.GetCategoriesList)
	apiPref.GET("/categories/:slug", categoryController.GetCategoryBySlug)
	apiPref.GET("/categories/:slug/products", productController.GetProductsByCategory)
	photosCache := &middleware.SafeCache{
		Data: make(map[string]*middleware.SimpleCacheItem),
	}

	apiPref.GET("/photos/:filename", middleware.SimpleCacheMiddleware(PhotoCacheDuration, photosCache), photosHandler())
	apiPref.GET("/homepage", homepageController.GetHomepageCategories)

	r.Static("/js", FrontendJSDir)
	r.Static("/css", FrontendCSSDir)

	r.StaticFile("/sitemap.xml", filepath.Join(FrontendDistDir, "sitemap.xml"))
	r.StaticFile("/robots.txt", filepath.Join(FrontendDistDir, "robots.txt"))
	r.StaticFile("/favicon.ico", filepath.Join(FrontendDistDir, "favicon.ico"))
	r.NoRoute(func(c *gin.Context) {
		path := FrontendDistDir + c.Request.URL.Path
		if filepath.Ext(path) != "" {
			c.Status(http.StatusNotFound)
			return
		}
		c.File(filepath.Join(FrontendDistDir, "index.html"))
	})

	return r
}

func photosHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := fmt.Sprintf("./photos/%s", filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			metrics.Error404Counter.WithLabelValues("404").Inc()
			api.SendError(c, http.StatusNotFound, "Файл не найден")
			return
		}

		c.Header("Cache-Control", "public, max-age=2592000, immutable")
		c.File(filePath)
	}
}
