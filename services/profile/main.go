package main

import (
	"github.com/dreadster3/pawcare/services/profile/api"
	docs "github.com/dreadster3/pawcare/services/profile/docs"
	"github.com/dreadster3/pawcare/services/profile/middleware"
	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func recovery(c *gin.Context, err any) {
	c.AbortWithStatusJSON(500, models.NewErrorResponseString(c, "Internal server error"))
}

func main() {
	godotenv.Load()

	router := gin.New()
	router.Use(gin.Logger(), gin.CustomRecovery(recovery), middleware.RequestIdMiddleware)

	api.RegisterRoutes(&router.RouterGroup)

	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Title = "Pet Profile Service"
		docs.SwaggerInfo.Description = "This is the Pet Profile Service"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Run(":8080")
}
