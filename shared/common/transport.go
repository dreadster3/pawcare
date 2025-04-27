package common

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dreadster3/pawcare/shared/utils"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

func KafkaDecodeJSONMessage[T any](_ context.Context, msg *message.Message) (interface{}, error) {
	var request T
	if err := json.Unmarshal(msg.Payload, &request); err != nil {
		return nil, err
	}
	return request, nil
}

func KafkaEncodeResponse(ctx context.Context, response interface{}) error {
	if e, ok := response.(endpoint.Failer); ok && e.Failed() != nil {
		// TODO: Change this
		return nil
	}
	return nil
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

func ErrorEncoder(err2Status func(error) int) func(_ context.Context, err error, w http.ResponseWriter) {
	return func(_ context.Context, err error, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(err2Status(err))
		json.NewEncoder(w).Encode(NewErrorResponse(err))
	}
}

func EncodeResponse(err2Status func(error) int) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if e, ok := response.(endpoint.Failer); ok && e.Failed() != nil {
			ErrorEncoder(err2Status)(ctx, e.Failed(), w)
			return nil
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(response)
	}
}

func DefaultErr2Status(err error) int {
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		return http.StatusBadRequest
	}

	switch err {
	case kitjwt.ErrTokenExpired, kitjwt.ErrTokenContextMissing, kitjwt.ErrTokenInvalid, kitjwt.ErrTokenMalformed, kitjwt.ErrTokenNotActive, jwt.ErrSignatureInvalid:
		return http.StatusUnauthorized
	case ErrAlreadyCreated:
		return http.StatusConflict
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
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
			requestId := ctx.Value(utils.RequestIdContextKey).(string)

			logger.Log("request_id", requestId, "method", method, "path", path, "client_ip", clientIp, "user_agent", userAgent, "content_length", contentLength, "msg", "Incoming request")

			return ctx
		}),

		httptransport.ServerFinalizer(func(ctx context.Context, code int, r *http.Request) {
			method := r.Method
			path := r.URL.Path
			clientIp := r.RemoteAddr
			requestId := ctx.Value(utils.RequestIdContextKey).(string)
			userAgent := r.UserAgent()

			logger.Log("request_id", requestId, "method", method, "path", path, "client_ip", clientIp, "user_agent", userAgent, "status_code", code, "msg", "Outgoing response")
		}),
	}
}

func GRPCLoggingServerOptions(logger log.Logger) []grpctransport.ServerOption {
	return []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			requestId := ctx.Value(utils.RequestIdContextKey).(string)
			logs := []interface{}{"request_id", requestId}
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
			requestId := ctx.Value(utils.RequestIdContextKey).(string)
			logs := []interface{}{"request_id", requestId}

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
