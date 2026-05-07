package users_postgres_repository

import (
	"context"
	"fmt"

	core_user_domain "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/domain/user"
)

func (p *UsersPostgresRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]core_user_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.GetTimeout())
	defer cancel()

	query := `
		SELECT id, version, full_name, password_hash, email, phone_number, telegram 
		FROM notifapp.users
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2;
	`

	rows, err := p.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}
	defer rows.Close()

	var userModel []UserModel
	for rows.Next() {
		var um UserModel

		err := rows.Scan(
			&um.ID,
			&um.Version,
			&um.FullName,
			&um.PasswordHash,
			&um.Email,
			&um.PhoneNumber,
			&um.Telegram,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		userModel = append(userModel, um)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("next row: %w", err)
	}

	var userDomain []core_user_domain.User
	for _, um := range userModel {
		userDomain = append(userDomain, core_user_domain.User{
			ID:           um.ID,
			Version:      um.Version,
			FullName:     um.FullName,
			PasswordHash: um.PasswordHash,
			Email:        um.Email,
			PhoneNumber:  um.PhoneNumber,
			Telegram:     um.Telegram,
		})
	}

	return userDomain, nil
}
