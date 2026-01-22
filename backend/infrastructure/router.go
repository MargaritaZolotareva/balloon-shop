package infrastructure

import (
	"backend/api"
	"backend/controllers"
	"backend/infrastructure/metrics"
	"backend/infrastructure/rabbitmq"
	"backend/repositories"
	"backend/services"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/openzipkin/zipkin-go"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func SetupRouter(db *gorm.DB, rmq *rabbitmq.RabbitMQ, tracer *zipkin.Tracer) *gin.Engine {
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

	r.Use(zipkinMiddleware(tracer))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.POST("/lead-form", messageController.SendLeadMessage)
	r.GET("/products/:id", productController.GetProduct)
	r.GET("/categories/:id", categoryController.GetCategory)
	r.GET("/categories/:id/products", productController.GetProductsByCategory)
	r.GET("/photos/:filename", photosHandler())

	return r
}

func zipkinMiddleware(tracer *zipkin.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Starting Zipkin span for %s %s", c.Request.Method, c.Request.URL.Path)

		span := tracer.StartSpan(c.Request.URL.Path)
		c.Set("zipkin-span", span)
		c.Request = c.Request.WithContext(zipkin.NewContext(c, span))
		log.Printf("Created Zipkin span for %s", c.Request.URL.Path)

		c.Next()

		span.Finish()
		log.Printf("Finished Zipkin span for %s %s", c.Request.Method, c.Request.URL.Path)
	}
}

func photosHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := fmt.Sprintf("./photos/%s", filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			sentry.CaptureException(err)
			metrics.Error404Counter.WithLabelValues("404").Inc()
			api.SendError(c, http.StatusNotFound, "Файл не найден")
			return
		}

		c.File(filePath)
	}
}
