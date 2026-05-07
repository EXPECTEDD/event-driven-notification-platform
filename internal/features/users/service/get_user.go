package users_service

import (
	"context"
	"fmt"

	core_http_utils "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/transport/http/utils"
)

type GetUserOutput struct {
	FullName    string
	Email       string
	PhoneNumber *string
	Telegram    *string
}

func (s *UserService) GetUser(
	ctx context.Context,
	id int,
) (GetUserOutput, error) {
	err := core_http_utils.ValidateID(id)
	if err != nil {
		return GetUserOutput{}, fmt.Errorf(
			"validate ID: %w",
			err,
		)
	}

	userDomain, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return GetUserOutput{}, fmt.Errorf(
			"get user from repository: %w",
			err,
		)
	}

	out := GetUserOutput{
		FullName:    userDomain.FullName,
		Email:       userDomain.Email,
		PhoneNumber: userDomain.PhoneNumber,
		Telegram:    userDomain.Telegram,
	}

	return out, nil
}
