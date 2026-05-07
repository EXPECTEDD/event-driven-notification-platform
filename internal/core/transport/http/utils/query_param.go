package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param=%s by key=%s not a valid integer: %w",
			param,
			key,
			core_error.ErrInvalidArgument,
		)
	}

	return &val, nil
}
