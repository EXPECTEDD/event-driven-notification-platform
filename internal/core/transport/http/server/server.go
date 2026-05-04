package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	core_http_middleware "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/middleware"
)

type HTTPServer struct {
	log *core_logger.Logger
	cfg Config
	mux *http.ServeMux
}

func NewHTTPServer(
	log *core_logger.Logger,
	cfg Config,
) *HTTPServer {
	mux := http.NewServeMux()

	return &HTTPServer{
		log: log,
		cfg: cfg,
		mux: mux,
	}
}

func (s *HTTPServer) Run(
	ctx context.Context,
	middlewares ...core_http_middleware.Middleware,
) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, middlewares...)

	server := http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port),
		Handler:      mux,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	chErr := make(chan error, 1)
	go func() {
		defer close(chErr)
		s.log.Info("starting HTTP server")
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				chErr <- err
			}
		}
	}()

	select {
	case err := <-chErr:
		if err != nil {
			return fmt.Errorf("listen and server HTTP server: %w", err)
		}
	case <-ctx.Done():
		s.log.Info("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			s.cfg.ShutdownTimeout,
		)

		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		s.log.Info("HTTP server stoped")
	}

	return nil
}

func (s *HTTPServer) SetRoutes(routes ...Route) {
	for _, r := range routes {
		s.mux.Handle(fmt.Sprintf("%s %s", r.Method, r.Path), r.Handler)
	}
}
