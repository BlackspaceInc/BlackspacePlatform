package api

import (
	"net/http"

	"k8s.io/klog/v2"
)

type LoggingMiddleware struct {
}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (m *LoggingMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		klog.Info(
			"request started",
			"proto", r.Proto,
			"uri", r.RequestURI,
			"method", r.Method,
			"remote", r.RemoteAddr,
			"user-agent", r.UserAgent(),
		)
		next.ServeHTTP(w, r)
	})
}
