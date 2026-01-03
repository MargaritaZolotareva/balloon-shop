package infrastructure

import (
	"backend/controllers"
	"backend/infrastructure/rabbitmq"
	"backend/repositories"
	"backend/services"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func SetupRouter(db *gorm.DB, rmq *rabbitmq.RabbitMQ) *gin.Engine {

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService()
	productController := controllers.NewProductController(db, productRepo, productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryController := controllers.NewCategoryController(db, categoryRepo)

	messageController := controllers.NewMessageController(rmq)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("VUE_APP_URL")},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	r.POST("/lead-form", messageController.SendLeadMessage)
	r.GET("/products/:id", productController.GetProduct)
	r.GET("/categories/:id", categoryController.GetCategory)
	r.GET("/categories/:id/products", productController.GetProductsByCategory)
	r.GET("/photos/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := fmt.Sprintf("./photos/%s", filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Файл не найден"})
			return
		}

		c.File(filePath)
	})

	return r
}
