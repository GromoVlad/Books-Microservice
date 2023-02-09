package routes

import (
	"fmt"
	"github.com/GromoVlad/go_microsrv_books/docs"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/createBook"
	createBookProtobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/createBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/deleteBook"
	deleteBookProtobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/deleteBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/findBook"
	findBookProtobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/findBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/listBook"
	listBookProtobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/listBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/controllers/updateBook"
	updateBookProtobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/updateBook/gRPC"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"google.golang.org/grpc"
	"log"
	"net"
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

	/** Подключение сервисов к gRPC */
	gRPC()

	/** Документация проекта */
	swaggerInfo(docs.SwaggerInfo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(os.Getenv("MICROSERVICE_PORT"))
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}

func gRPC() {
	listener, err := net.Listen("tcp", fmt.Sprintf(os.Getenv("MICROSERVICE_BOOKS_GRPC_ADDRESS")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	findBookProtobuf.RegisterFindBookServer(grpcServer, findBook.NewServer())
	listBookProtobuf.RegisterListBookServer(grpcServer, listBook.NewServer())
	createBookProtobuf.RegisterCreateBookServer(grpcServer, createBook.NewServer())
	updateBookProtobuf.RegisterUpdateBookServer(grpcServer, updateBook.NewServer())
	deleteBookProtobuf.RegisterDeleteBookServer(grpcServer, deleteBook.NewServer())
	grpcServer.Serve(listener)
}

func swaggerInfo(swaggerInfo *swag.Spec) {
	swaggerInfo.Title = os.Getenv("MICROSERVICE_TITLE")
	swaggerInfo.Description = os.Getenv("MICROSERVICE_DESCRIPTION")
	swaggerInfo.Version = os.Getenv("MICROSERVICE_VERSION")
	swaggerInfo.Host = os.Getenv("MICROSERVICE_HOST")
	swaggerInfo.BasePath = os.Getenv("MICROSERVICE_BASE_PATH")
	swaggerInfo.Schemes = []string{"http"}
}
