package users_service

import (
	"context"
	"fmt"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int,
) error {
	if id <= 0 {
		return fmt.Errorf("invalid `id` value: %w",
			core_error.ErrInvalidArgument,
		)
	}

	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("delete user from repository: %w",
			err,
		)
	}

	return nil
}
