package infrastructure

import (
	"backend/api"
	"backend/controllers"
	"backend/infrastructure/metrics"
	"backend/repositories"
	"backend/services"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
	"net/http"
	"os"
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

	apiPref := r.Group("/api")
	apiPref.GET("/metrics", gin.WrapH(promhttp.Handler()))
	apiPref.POST("/lead-form", messageController.SendLeadMessage)
	apiPref.GET("/products/:slug", productController.GetProductBySlug)
	apiPref.GET("/categories", categoryController.GetCategoriesList)
	apiPref.GET("/categories/:slug", categoryController.GetCategoryBySlug)
	apiPref.GET("/categories/:slug/products", productController.GetProductsByCategory)
	apiPref.GET("/photos/:filename", photosHandler())
	apiPref.GET("/homepage", homepageController.GetHomepageCategories)

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

		c.File(filePath)
	}
}
