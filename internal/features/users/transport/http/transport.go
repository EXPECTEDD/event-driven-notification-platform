package users_transport_http

import (
	"context"
	"net/http"

	core_http_server "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/server"
	users_service "github.com/EXPECTEDD/event-driven-notification-platform/internal/features/users/service"
)

type UserService interface {
	RegisterUser(
		ctx context.Context,
		input users_service.RegisterUserInput,
	) (users_service.RegisterUserOutput, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error
}

type UserHTTPHandler struct {
	userService UserService
}

func NewUserHTTPHandler(
	userService UserService,
) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) GetRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/user/registration",
			Method:  http.MethodPost,
			Handler: h.RegisterUser,
		}, {
			Path:    "/user/{id}",
			Method:  http.MethodDelete,
			Handler: h.DeleteUser,
		},
	}
}
