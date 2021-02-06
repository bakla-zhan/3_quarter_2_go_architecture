package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	CtxRequestIDKey = "X-Request-Id"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.Context().Value(CtxRequestIDKey)
		if v == nil {
			v = uuid.New().String()
			ctxWithRID := context.WithValue(r.Context(), CtxRequestIDKey, v)
			r = r.WithContext(ctxWithRID)
		}
		next.ServeHTTP(w, r)
	})
}
