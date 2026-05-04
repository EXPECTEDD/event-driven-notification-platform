package users_service

import (
	"context"
	"fmt"

	core_user_domain "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/domain/user"
	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

type RegisterUserInput struct {
	FullName    string
	Password    string
	Email       string
	PhoneNumber *string
	Telegram    *string
}

type RegisterUserOutput struct {
	FullName    string
	Email       string
	PhoneNumber *string
	Telegram    *string
}

func (s *UserService) RegisterUser(
	ctx context.Context,
	input RegisterUserInput,
) (RegisterUserOutput, error) {
	if err := validatePassword(input.Password); err != nil {
		return RegisterUserOutput{}, fmt.Errorf(
			"validate password: %w",
			err,
		)
	}

	passwordHash, err := s.hasher.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		return RegisterUserOutput{}, fmt.Errorf(
			"generate password hash: %w",
			err,
		)
	}

	userDomain, err := core_user_domain.NewUninitializedUser(
		input.FullName,
		string(passwordHash),
		input.Email,
		input.PhoneNumber,
		input.Telegram,
	)
	if err != nil {
		return RegisterUserOutput{}, fmt.Errorf(
			"create user domain: %w",
			err,
		)
	}

	user, err := s.userRepository.SaveUser(ctx, userDomain)
	if err != nil {
		return RegisterUserOutput{}, fmt.Errorf(
			"save user to repository: %w",
			err,
		)
	}

	return RegisterUserOutput{
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Telegram:    user.Telegram,
	}, nil
}

func validatePassword(password string) error {
	passwordLen := len([]rune(password))
	if passwordLen < 5 || passwordLen > 100 {
		return fmt.Errorf(
			"invalid `Password` len: %d: %w",
			passwordLen,
			core_error.ErrInvalidArgument,
		)
	}
	return nil
}
