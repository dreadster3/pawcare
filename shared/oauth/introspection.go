package oauth

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func introspectToken(introspectionEndpoint string, clientId string, clientSecret string, token string) (jwt.Claims, error) {
	data := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"token":         {token},
	}

	req, err := http.NewRequest("POST", introspectionEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	transport := http.DefaultTransport
	transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	active, ok := result["active"].(bool)
	if !ok {
		return nil, fmt.Errorf("no active field")
	}

	if !active {
		return nil, fmt.Errorf("token is inactive")
	}

	claims := &jwt.StandardClaims{}
	claims.Subject = result["sub"].(string)
	claims.ExpiresAt = int64(result["exp"].(float64))
	claims.IssuedAt = int64(result["iat"].(float64))
	return claims, nil
}

func NewIntrospectionMiddleware(introspectionEndpoint string, clientId string, clientSecret string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// tokenString is stored in the context from the transport handlers.
			tokenString, ok := ctx.Value(kitjwt.JWTContextKey).(string)
			if !ok {
				return nil, kitjwt.ErrTokenContextMissing
			}

			claims, err := introspectToken(introspectionEndpoint, clientId, clientSecret, tokenString)
			if err != nil {
				return nil, err
			}

			if err := claims.Valid(); err != nil {
				return nil, err
			}

			ctx = context.WithValue(ctx, kitjwt.JWTClaimsContextKey, claims)

			return next(ctx, request)
		}
	}
}

func NewConfiguredIntrospectionMiddleware(viper viper.Viper) endpoint.Middleware {
	introspectionEndpoint := viper.GetString(common.IntrospectionEndpointKey)
	clientId := viper.GetString(common.ClientIdKey)
	clientSecret := viper.GetString(common.ClientSecretKey)
	return NewIntrospectionMiddleware(introspectionEndpoint, clientId, clientSecret)
}
