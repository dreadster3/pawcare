package http

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrNoCookieProvided   = errors.New("no cookie provided")
	ErrNoHeaderProvided   = errors.New("no header provided")
	ErrInvalidTokenFormat = errors.New("invalid token")
	ErrNoTokenFound       = errors.New("no token found")
)

type ServiceContext struct {
	Host string
	c    *gin.Context
}

func NewServiceContext(host string, c *gin.Context) *ServiceContext {
	return &ServiceContext{
		Host: host,
		c:    c,
	}
}

func getTokenFromCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		return "", ErrNoCookieProvided
	}

	return cookie, nil
}

func getTokenFromHeader(c *gin.Context) (string, error) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		return "", ErrNoHeaderProvided
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", ErrInvalidTokenFormat
	}

	return strings.TrimPrefix(authorization, "Bearer "), nil
}

func (s *ServiceContext) BearerToken() string {
	token, err := getTokenFromHeader(s.c)
	if err != nil {
		token, err = getTokenFromCookie(s.c)
		if err != nil {
			return ""
		}
	}

	return token
}
