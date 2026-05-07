package users_service

import (
	"context"
	"fmt"

	core_http_utils "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/utils"
)

type GetUsersOutput struct {
	FullName    string
	Email       string
	PhoneNumber *string
	Telegram    *string
}

func (s *UserService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]GetUsersOutput, error) {
	if limit != nil {
		if err := core_http_utils.ValidateLimit(*limit); err != nil {
			return []GetUsersOutput{}, fmt.Errorf(
				"validate limit: %w",
				err,
			)
		}
	}
	if offset != nil {
		if err := core_http_utils.ValidateOffset(*offset); err != nil {
			return []GetUsersOutput{}, fmt.Errorf(
				"validate offset: %w",
				err,
			)
		}
	}

	users, err := s.userRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return []GetUsersOutput{}, fmt.Errorf(
			"get users from repository: %w",
			err,
		)
	}

	var out []GetUsersOutput
	for _, u := range users {
		out = append(out, GetUsersOutput{
			FullName:    u.FullName,
			Email:       u.Email,
			PhoneNumber: u.PhoneNumber,
			Telegram:    u.Telegram,
		})
	}

	return out, nil
}
