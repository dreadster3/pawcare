package common

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dreadster3/pawcare/shared/utils"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func DecodeNoBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodePathParameters[T any](_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	result, err := utils.MapToStruct[T](vars)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DecodeJSONRequest[T any](_ context.Context, r *http.Request) (interface{}, error) {
	var request T
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func GRPCDecodeNoBody(_ context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}

func GRPCDecodeToObject[T any](_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(T)
	if !ok {
		return nil, ErrCastRequest
	}

	return req, nil
}

type errorer interface {
	error() error
}

func EncodeResponse(encodeError func(context.Context, error, http.ResponseWriter)) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if e, ok := response.(errorer); ok && e.error() != nil {
			encodeError(ctx, e.error(), w)
			return nil
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(response)
	}
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch err {
	case kitjwt.ErrTokenExpired, kitjwt.ErrTokenContextMissing, kitjwt.ErrTokenInvalid, kitjwt.ErrTokenMalformed, kitjwt.ErrTokenNotActive, jwt.ErrSignatureInvalid:
		w.WriteHeader(http.StatusUnauthorized)
	case ErrAlreadyCreated:
		w.WriteHeader(http.StatusConflict)
	case ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HTTPLoggingServerOptions(logger log.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			clientIp := r.RemoteAddr
			method := r.Method
			path := r.URL.Path
			contentLength := r.ContentLength
			userAgent := r.UserAgent()

			logger.Log("method", method, "path", path, "client_ip", clientIp, "user_agent", userAgent, "content_length", contentLength, "msg", "Incoming request")

			return ctx
		}),

		httptransport.ServerFinalizer(func(ctx context.Context, code int, r *http.Request) {
			method := r.Method
			path := r.URL.Path
			clientIp := r.RemoteAddr
			userAgent := r.UserAgent()

			logger.Log("method", method, "path", path, "client_ip", clientIp, "user_agent", userAgent, "status_code", code, "msg", "Outgoing response")
		}),
	}
}

func GRPCLoggingServerOptions(logger log.Logger) []grpctransport.ServerOption {
	return []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			logs := []interface{}{}
			if userAgents, ok := md["user-agent"]; ok {
				logs = append(logs, "user_agent", strings.Join(userAgents, ";"))
			}

			if p, ok := peer.FromContext(ctx); ok && p != nil {
				logs = append(logs, "client_ip", p.Addr)
			}

			logs = append(logs, "msg", "Incoming request")

			logger.Log(logs...)

			return ctx
		}),
		grpctransport.ServerFinalizer(func(ctx context.Context, err error) {
			logs := []interface{}{}

			if md, ok := metadata.FromIncomingContext(ctx); ok {
				if userAgents, ok := md["user-agent"]; ok {
					logs = append(logs, "user_agent", strings.Join(userAgents, ";"))
				}
			}

			if p, ok := peer.FromContext(ctx); ok && p != nil {
				logs = append(logs, "client_ip", p.Addr)
			}
			logs = append(logs, "msg", "Outgoing response")

			logger.Log(logs...)
		}),
	}
}
