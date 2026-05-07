package users_transport_http

import (
	"net/http"

	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	core_http_response "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/response"
	core_http_utils "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/utils"
)

type GetUsersResponse struct {
	FullName    string
	Email       string
	PhoneNumber *string
	Telegram    *string
}

func (h *UserHTTPHandler) GetUsers(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContextOrPanic(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		responseHandler.ErrorResponse("get `limit` query param", err)
		return
	}

	if limit != nil {
		if err := core_http_utils.ValidateLimit(*limit); err != nil {
			responseHandler.ErrorResponse("validate limit", err)
			return
		}
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		responseHandler.ErrorResponse("get `offset` query param", err)
		return
	}

	if offset != nil {
		if err := core_http_utils.ValidateOffset(*offset); err != nil {
			responseHandler.ErrorResponse("validate offset", err)
			return
		}
	}

	out, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse("get users", err)
		return
	}

	var response []GetUsersResponse
	for _, u := range out {
		response = append(response, GetUsersResponse{
			FullName:    u.FullName,
			Email:       u.Email,
			PhoneNumber: u.PhoneNumber,
			Telegram:    u.Telegram,
		})
	}

	responseHandler.SendResponse(response, http.StatusOK)
}
