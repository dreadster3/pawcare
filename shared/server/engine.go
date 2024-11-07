package server

import (
	"github.com/dreadster3/pawcare/shared/middleware"
	"github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
)

func recovery(c *gin.Context, err any) {
	c.AbortWithStatusJSON(500, models.NewErrorResponseString(c, "Internal server error"))
}

func NewDefaultEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.CustomRecovery(recovery), middleware.RequestIdMiddleware)
	return router
}
