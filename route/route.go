package route

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"

	"rishabh/rest-api/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Api Prefix
	apiPrefix := os.Getenv("API_PREFIX")
	// version 1
	apiVersion := os.Getenv("API_VERSION")

	apiV1 := r.Group(apiPrefix / apiVersion)
	// Apply the middleware to the router (works with groups too)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	user := apiV1.Group("/users")
	{
		user.GET("/", controller.GetUsers)
		user.POST("/", controller.CreateUser)
		user.GET("/:id", controller.GetUserByID)
	}
	return r
}
