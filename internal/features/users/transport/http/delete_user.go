package users_transport_http

import (
	"fmt"
	"net/http"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
	core_logger "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/logger"
	core_http_response "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/response"
	core_http_utils "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/utils"
)

func (h *UserHTTPHandler) DeleteUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContextOrPanic(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	id, err := core_http_utils.GetPathValueInt(r, "id")
	if err != nil {
		responseHandler.ErrorResponse("get path value", err)
		return
	}

	err = validateID(id)
	if err != nil {
		responseHandler.ErrorResponse("validate id", err)
		return
	}

	err = h.userService.DeleteUser(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse("delete user", err)
		return
	}

	responseHandler.NoContentResponse()
}

func validateID(id int) error {
	if id <= 0 {
		return fmt.Errorf(
			"invalid `id` value: %d: %w",
			id,
			core_error.ErrInvalidArgument,
		)
	}
	return nil
}
