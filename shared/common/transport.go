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
	kittransport "github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func HTTPLoggingServerOptions(logger *zap.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			clientIp := r.RemoteAddr
			method := r.Method
			path := r.URL.Path
			contentLength := r.ContentLength
			userAgent := r.UserAgent()
			requestId := ctx.Value(utils.RequestIdContextKey).(string)

			logger.
				Info("Incoming request",
					zap.String("request_id", requestId),
					zap.String("method", method),
					zap.String("path", path),
					zap.String("client_ip", clientIp),
					zap.String("user_agent", userAgent),
					zap.Int64("content_length", contentLength),
				)

			return ctx
		}),

		httptransport.ServerFinalizer(func(ctx context.Context, code int, r *http.Request) {
			method := r.Method
			path := r.URL.Path
			clientIp := r.RemoteAddr
			requestId := ctx.Value(utils.RequestIdContextKey).(string)
			userAgent := r.UserAgent()

			logger.
				Info("Outgoing response",
					zap.String("request_id", requestId),
					zap.String("method", method),
					zap.String("path", path),
					zap.String("client_ip", clientIp),
					zap.String("user_agent", userAgent),
					zap.Int("status_code", code),
				)
		}),
	}
}

func GRPCLoggingServerOptions(logger *zap.Logger) []grpctransport.ServerOption {
	return []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			requestId := ctx.Value(utils.RequestIdContextKey).(string)
			logs := []zapcore.Field{zap.String("request_id", requestId)}
			if userAgents, ok := md["user-agent"]; ok {
				logs = append(logs, zap.String("user_agent", strings.Join(userAgents, ";")))
			}

			if p, ok := peer.FromContext(ctx); ok && p != nil {
				logs = append(logs, zap.Stringer("client_ip", p.Addr))
			}

			logger.Info("Incoming request", logs...)

			return ctx
		}),
		grpctransport.ServerFinalizer(func(ctx context.Context, err error) {
			requestId := ctx.Value(utils.RequestIdContextKey).(string)
			fields := []zapcore.Field{zap.String("request_id", requestId)}

			if md, ok := metadata.FromIncomingContext(ctx); ok {
				if userAgents, ok := md["user-agent"]; ok {
					fields = append(fields, zap.String("user_agent", strings.Join(userAgents, ";")))
				}
			}

			if p, ok := peer.FromContext(ctx); ok && p != nil {
				fields = append(fields, zap.Stringer("client_ip", p.Addr))
			}

			logger.Info("Outgoing response", fields...)
		}),
	}
}

type logErrorHandler struct {
	logger *zap.Logger
}

func NewLogErrorHandler(logger *zap.Logger) kittransport.ErrorHandler {
	return &logErrorHandler{logger}
}

func (h *logErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Error(err.Error())
}
