package http

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
)

func BearerToken(token string) func(*runtime.ClientOperation) {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = httptransport.BearerToken(token)
	}
}
