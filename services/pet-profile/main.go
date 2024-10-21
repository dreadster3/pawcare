package main

import (
	"github.com/dreadster3/pawcare/services/pet-profile/api"
	docs "github.com/dreadster3/pawcare/services/pet-profile/docs"
	"github.com/dreadster3/pawcare/services/pet-profile/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	godotenv.Load()

	router := gin.Default()
	router.Use(middleware.RequestIdMiddleware)

	api.RegisterRoutes(&router.RouterGroup)

	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Title = "Pet Profile Service"
		docs.SwaggerInfo.Description = "This is the Pet Profile Service"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Run(":8080")
}
