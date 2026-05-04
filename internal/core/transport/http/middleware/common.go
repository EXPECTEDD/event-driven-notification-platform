package core_http_middleware

import (
	"net/http"
	"time"

	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	core_http_response "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDStr = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDStr)
			if requestID == "" {
				requestID = uuid.NewString()
				r.Header.Set(requestIDStr, requestID)
			}

			w.Header().Set(requestIDStr, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDStr)

			l := log.With(
				zap.String("reqiest_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := core_logger.ToContext(
				r.Context(),
				l,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContextOrPanic(ctx)

			timestamp := time.Now()

			log.Info(
				"--- incoming HTTP request",
				zap.String("method", r.Method),
				zap.Time("time", timestamp.UTC()),
			)

			rw := core_http_response.NewResponseWriter(w)

			next.ServeHTTP(rw, r)

			log.Info(
				"--- done HTTP request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Since(timestamp)),
			)
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContextOrPanic(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(
				w,
				log,
			)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HTTP request got unexpected panic",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
