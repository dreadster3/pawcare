package middleware

import (
	"errors"
	"strings"

	"github.com/dreadster3/pawcare/services/profile/env"
	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/gin-gonic/gin"
)

var (
	NotAuthorized = errors.New("not authorized")
)

func getTokenFromCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		return "", NotAuthorized
	}

	return cookie, nil
}

func getTokenFromHeader(c *gin.Context) (string, error) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		return "", NotAuthorized
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", NotAuthorized
	}

	return strings.TrimPrefix(authorization, "Bearer "), nil
}

func JwtAuth(env *env.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeader(c)
		if err != nil {
			token, err = getTokenFromCookie(c)
			if err != nil {
				c.AbortWithStatusJSON(401, models.NewErrorResponse(c, err))
				return
			}
		}

		claims, err := env.Services.Auth.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, models.NewErrorResponse(c, NotAuthorized))
			return
		}

		userId, err := claims.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(401, models.NewErrorResponse(c, NotAuthorized))
			return
		}

		c.Set("user_id", userId)

		c.Next()
	}
}
