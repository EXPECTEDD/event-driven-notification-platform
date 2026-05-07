package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

var (
	invalidInt = -1
)

func GetIntPathValue(
	r *http.Request,
	key string,
) (int, error) {
	path := r.PathValue(key)
	if path == "" {
		return invalidInt, fmt.Errorf(
			"empty `%s` path value: %w",
			key,
			core_error.ErrInvalidArgument,
		)
	}

	val, err := strconv.Atoi(path)
	if err != nil {
		return invalidInt, fmt.Errorf(
			"path=%s by key=%s not a valid integer: %w",
			path,
			key,
			core_error.ErrInvalidArgument,
		)
	}

	return val, nil
}
