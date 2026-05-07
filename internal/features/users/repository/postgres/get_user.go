package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	core_user_domain "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/domain/user"
	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
	"github.com/jackc/pgx/v5"
)

func (p *UsersPostgresRepository) GetUser(
	ctx context.Context,
	id int,
) (core_user_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.GetTimeout())
	defer cancel()

	query := `
		SELECT id, version, full_name, password_hash, email, phone_number, telegram 
		FROM notifapp.users
		WHERE id = $1;
	`

	row := p.pool.QueryRow(ctx, query, id)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PasswordHash,
		&userModel.Email,
		&userModel.PhoneNumber,
		&userModel.Telegram,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_user_domain.User{}, core_error.ErrNotFound
		}
		return core_user_domain.User{}, fmt.Errorf(
			"scan row: %w",
			err,
		)
	}

	userDomain, err := core_user_domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PasswordHash,
		userModel.Email,
		userModel.PhoneNumber,
		userModel.Telegram,
	)
	if err != nil {
		return core_user_domain.User{}, fmt.Errorf(
			"create user domain: %w",
			err,
		)
	}

	return userDomain, nil
}
