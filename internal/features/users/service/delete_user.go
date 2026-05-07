package users_service

import (
	"context"
	"fmt"

	core_http_utils "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/utils"
)

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int,
) error {
	err := core_http_utils.ValidateID(id)
	if err != nil {
		return fmt.Errorf("validate ID: %w", err)
	}

	err = s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("delete user from repository: %w",
			err,
		)
	}

	return nil
}
