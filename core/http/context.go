package core_http

import (
	"context"
	"net/http"
)

type contextKey string

func buildContext(req *http.Request, paramKeys, paramValues []string) *http.Request {
	ctx := req.Context()

	for i := 0; i < len(paramKeys); i++ {
		ctx = context.WithValue(ctx, contextKey(paramKeys[i]), paramValues[i])
	}

	return req.WithContext(ctx)
}
