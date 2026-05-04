package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	core_user_domain "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/domain/user"
	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	UniqueViolation = "23505"
)

func (p *UsersPostgresRepository) SaveUser(
	ctx context.Context,
	user core_user_domain.User,
) (core_user_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.GetTimeout())
	defer cancel()

	query := `
		INSERT INTO notifapp.users (full_name, password_hash, email, phone_number, telegram)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, version, full_name, password_hash, email, phone_number, telegram;
	`

	row := p.pool.QueryRow(ctx, query,
		user.FullName,
		user.PasswordHash,
		user.Email,
		user.PhoneNumber,
		user.Telegram,
	)

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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == UniqueViolation {
				return core_user_domain.User{},
					fmt.Errorf("scan row: %w: %v",
						core_error.ErrConflict,
						err,
					)
			}
		}
		return core_user_domain.User{},
			fmt.Errorf("scan row: %w", err)
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
		return core_user_domain.User{},
			fmt.Errorf("create user domain: %w", err)
	}

	return userDomain, nil
}
