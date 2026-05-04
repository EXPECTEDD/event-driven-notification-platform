package users_service

import (
	"context"

	core_user_domain "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/domain/user"
)

type Hasher interface {
	GenerateFromPassword(
		[]byte,
	) ([]byte, error)
}

type UserRepository interface {
	SaveUser(
		context.Context,
		core_user_domain.User,
	) (core_user_domain.User, error)
}

type UserService struct {
	hasher         Hasher
	userRepository UserRepository
}

func NewUserService(
	hasher Hasher,
	userRepository UserRepository,
) *UserService {
	return &UserService{
		hasher:         hasher,
		userRepository: userRepository,
	}
}
