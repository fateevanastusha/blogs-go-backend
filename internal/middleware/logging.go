package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/fateevanastusha/blogs-go-backend/internal/logger"
)

// statusRecorder запоминает статус ответа: из http.ResponseWriter его не достать
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func Logging(base *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l := base.With(
				"request_id", RequestIDFromContext(r.Context()),
				"method", r.Method,
				"path", r.URL.Path,
			)
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

			l.Info("request started")

			next.ServeHTTP(rec, r.WithContext(logger.WithContext(r.Context(), l)))

			level := slog.LevelInfo
			if rec.status >= 500 {
				level = slog.LevelError
			}
			l.Log(r.Context(), level, "request completed",
				"status", rec.status,
				"duration_ms", time.Since(start).Milliseconds(),
			)
		})
	}
}
