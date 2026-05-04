package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

var (
	invalidID = -1
)

func GetPathValueInt(
	r *http.Request,
	val string,
) (int, error) {
	idStr := r.PathValue(val)
	if idStr == "" {
		return invalidID, fmt.Errorf(
			"empty `%s` path value: %w",
			val,
			core_error.ErrInvalidArgument,
		)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return invalidID, fmt.Errorf(
			"invalid '%s' type: %w",
			val,
			core_error.ErrInvalidArgument,
		)
	}

	return id, nil
}
