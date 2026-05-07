package core_http_utils

import (
	"fmt"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

func ValidateID(id int) error {
	if id <= 0 {
		return fmt.Errorf(
			"invalid `id` value: %d: %w",
			id,
			core_error.ErrInvalidArgument,
		)
	}
	return nil
}

func ValidateLimit(l int) error {
	if l < 0 {
		return fmt.Errorf(
			"invalid `limit` value: %w",
			core_error.ErrInvalidArgument,
		)
	}
	return nil
}

func ValidateOffset(o int) error {
	if o < 0 {
		return fmt.Errorf(
			"invalid `offset` value: %w",
			core_error.ErrInvalidArgument,
		)
	}
	return nil
}
