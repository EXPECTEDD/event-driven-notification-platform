package users_transport_http

import (
	"net/http"

	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	core_http_request "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/request"
	core_http_response "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/response"
	users_service "github.com/EXPECTEDD/event-driven-notification-platform/internal/features/users/service"
)

type RegisterUserRequest struct {
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"`
	Password    string  `json:"password"     validate:"required,min=5,max=100"`
	Email       string  `json:"email"        validate:"required,min=5,max=100,contains=@,contains=."`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
	Telegram    *string `json:"telegram"     validate:"omitempty,min=6,max=33,startswith=@"`
}

type RegisterUserResponse struct {
	FullName    string  `json:"full_name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Telegram    *string `json:"telegram"`
}

func (h *UserHTTPHandler) RegisterUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContextOrPanic(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	var req RegisterUserRequest
	if err := core_http_request.DecodeAndValidate(r, &req); err != nil {
		responseHandler.ErrorResponse("decode and validate request", err)
		return
	}

	out, err := h.userService.RegisterUser(ctx, users_service.RegisterUserInput{
		FullName:    req.FullName,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Telegram:    req.Telegram,
	})
	if err != nil {
		responseHandler.ErrorResponse("register user", err)
		return
	}

	response := RegisterUserResponse{
		FullName:    out.FullName,
		Email:       out.Email,
		PhoneNumber: out.PhoneNumber,
		Telegram:    out.Telegram,
	}

	responseHandler.SendResponse(response, http.StatusCreated)
}
