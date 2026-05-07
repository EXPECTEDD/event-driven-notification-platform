package users_postgres_repository

import (
	"context"
	"fmt"

	core_error "github.com/EXPECTEDD/event-driven-notification-platform/internal/core/error"
)

func (p *UsersPostgresRepository) DeleteUser(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, p.pool.GetTimeout())
	defer cancel()

	query := `
		DELETE FROM notifapp.users
		WHERE id = $1;
	`

	tag, err := p.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user with id=%d not found: %w", id, core_error.ErrNotFound)
	}

	return nil
}
