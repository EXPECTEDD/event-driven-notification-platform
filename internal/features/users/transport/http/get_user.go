package users_transport_http

import (
	"net/http"

	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	core_http_response "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/response"
	core_http_utils "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/utils"
)

type GetUserResponse struct {
	FullName    string
	Email       string
	PhoneNumber *string
	Telegram    *string
}

func (h *UserHTTPHandler) GetUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContextOrPanic(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	id, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse("get path value", err)
		return
	}

	err = core_http_utils.ValidateID(id)
	if err != nil {
		responseHandler.ErrorResponse("validate ID", err)
		return
	}

	out, err := h.userService.GetUser(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse("get user", err)
		return
	}

	response := GetUserResponse{
		FullName:    out.FullName,
		Email:       out.Email,
		PhoneNumber: out.PhoneNumber,
		Telegram:    out.Telegram,
	}

	responseHandler.SendResponse(response, http.StatusOK)
}
