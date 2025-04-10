package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func UUID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// add uuid to response header
		requestId := uuid.New().String()

		// using context append value
		ctx := context.WithValue(request.Context(), "X-Request-Id", requestId)

		// call next handler in chain
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
