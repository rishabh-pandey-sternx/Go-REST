package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"rishabh/rest-api/config"
	_ "rishabh/rest-api/docs"
	"rishabh/rest-api/model"
	"rishabh/rest-api/route"
)

// @title GO REST API documentation
// @version 1.0.0
// @description This is a documentation for REST API.
// @termsOfService http://swagger.io/terms/

// @contact.name Rishabh Pandey
// @contact.url https://geekrishabh.in
// @contact.email geekrishabh@gmail.com

// @host localhost:3000
// @BasePath /api/v1

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.DB = config.SetupDatabase()
	config.DB.AutoMigrate(&model.User{})

	r := route.SetupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
