package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	core_postgres_pool "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/middleware"
	core_http_server "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/server"
	users_hasher_repository "github.com/EXPECTEDD/event-driven-notification-platform/internal/features/users/repository/hasher"
	users_postgres_repository "github.com/EXPECTEDD/event-driven-notification-platform/internal/features/users/repository/postgres"
	users_service "github.com/EXPECTEDD/event-driven-notification-platform/internal/features/users/service"
	users_transport_http "github.com/EXPECTEDD/event-driven-notification-platform/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	logger := core_logger.NewLoggerMust(core_logger.NewConfigMust())
	defer logger.Close()

	postgresPool, err := core_postgres_pool.NewConnectionPool(
		core_postgres_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Error("create connection pool", zap.Error(err))
		os.Exit(1)
	}

	usersHashRepository := users_hasher_repository.NewUsersHasherRepository()
	usersPostgresRepository := users_postgres_repository.NewUsersPostgresRepository(postgresPool)

	usersService := users_service.NewUserService(usersHashRepository, usersPostgresRepository)

	usersHTTPHandler := users_transport_http.NewUserHTTPHandler(usersService)

	httpServer := core_http_server.NewHTTPServer(logger, core_http_server.NewConfigMust())

	httpServer.SetRoutes(usersHTTPHandler.GetRoutes()...)

	err = httpServer.Run(
		ctx,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	if err != nil {
		logger.Error("run http server", zap.Error(err))
		os.Exit(1)
	}
}
