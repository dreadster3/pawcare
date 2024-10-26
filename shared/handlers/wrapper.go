package handlers

import (
	"github.com/dreadster3/pawcare/shared/env"
	"github.com/gin-gonic/gin"
)

type HandlerFuncWithEnv[T env.IServiceContainer] func(env env.IEnvironment[T], c *gin.Context)

func WrapperEnv[T env.IServiceContainer, E env.IEnvironment[T]](env E, handler func(env E, c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(env, c)
	}
}
