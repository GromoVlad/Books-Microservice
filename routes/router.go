package routes

import (
	"fmt"
	"github.com/GromoVlad/go_microsrv_books/docs"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/createBook"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/deleteBook"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/findBook"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/listBook"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/updateBook"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"log"
	"os"
)

func Run() {
	router := gin.New()

	/** Восстановление после ошибки */
	router.Use(gin.Recovery())
	/** Подгружаем данные из .env */
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки переменных из .env: %s", err.Error())
	}

	/** Роуты */
	api(router)

	/** Документация проекта */
	swaggerInfo(docs.SwaggerInfo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(os.Getenv("MICROSERVICE_PORT"))
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}

// api - Роуты
func api(router *gin.Engine) {
	router.GET("/book/:bookId", findBook.Endpoint)
	router.GET("/book/list", listBook.Endpoint)
	router.POST("/book", createBook.Endpoint)
	router.PUT("/book/:bookId", updateBook.Endpoint)
	router.DELETE("/book/:bookId", deleteBook.Endpoint)
}

func swaggerInfo(swaggerInfo *swag.Spec) {
	swaggerInfo.Title = os.Getenv("MICROSERVICE_TITLE")
	swaggerInfo.Description = os.Getenv("MICROSERVICE_DESCRIPTION")
	swaggerInfo.Version = os.Getenv("MICROSERVICE_VERSION")
	swaggerInfo.Host = os.Getenv("MICROSERVICE_HOST")
	swaggerInfo.BasePath = os.Getenv("MICROSERVICE_BASE_PATH")
	swaggerInfo.Schemes = []string{"http"}
}
