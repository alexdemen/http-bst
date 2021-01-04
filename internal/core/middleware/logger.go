package middleware

import (
	"github.com/alexdemen/http-bst/internal/core/log"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type StructuredLogger struct {
	Logger *log.Logger
}

func NewStructuredLogger(logger *log.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(logger)
}
