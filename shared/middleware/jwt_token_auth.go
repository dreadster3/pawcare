package middleware

import (
	"errors"
	"strings"

	"github.com/dreadster3/pawcare/shared/env"
	"github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
)

var (
	ErrNotAuthorized = errors.New("not authorized")
)

func getTokenFromCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		return "", ErrNotAuthorized
	}

	return cookie, nil
}

func getTokenFromHeader(c *gin.Context) (string, error) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		return "", ErrNotAuthorized
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", ErrNotAuthorized
	}

	return strings.TrimPrefix(authorization, "Bearer "), nil
}

func JwtAuth[T env.IServiceContainer, E env.IEnvironment[T]](env E) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeader(c)
		if err != nil {
			token, err = getTokenFromCookie(c)
			if err != nil {
				c.AbortWithStatusJSON(401, models.NewErrorResponse(c, err))
				return
			}
		}

		jwtToken, err := env.Services().Auth().VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, models.NewErrorResponse(c, ErrNotAuthorized))
			return
		}

		userId, err := jwtToken.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(401, models.NewErrorResponse(c, ErrNotAuthorized))
			return
		}

		c.Set("user_id", userId)

		c.Next()
	}
}
