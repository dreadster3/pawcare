package handlers

import (
	"errors"

	"github.com/dreadster3/pawcare/shared/env"
	"github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
)

type EnvFactory[T env.IServiceContainer, E env.IEnvironment[T]] func(c *gin.Context) (E, error)

type HandlerFuncWithEnv[T env.IServiceContainer, E env.IEnvironment[T]] func(env E, c *gin.Context)

func WrapperEnv[T env.IServiceContainer, E env.IEnvironment[T]](factory EnvFactory[T, E], handler HandlerFuncWithEnv[T, E]) gin.HandlerFunc {
	return func(c *gin.Context) {
		env, err := factory(c)
		if err != nil {
			c.Error(errors.New("Failed to create environment"))
			c.Error(err)
			c.AbortWithStatusJSON(500, models.NewErrorResponseString(c, "Internal server error"))
			return
		}

		handler(env, c)
	}
}
