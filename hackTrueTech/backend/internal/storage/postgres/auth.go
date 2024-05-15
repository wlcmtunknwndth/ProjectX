package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/wlcmtunknwndth/hackBPA/internal/auth"
	"time"
)

func (s *Storage) GetPassword(ctx context.Context, username string) (string, error) {
	const op = "storage.postgres.auth.GetPassword"

	newCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	var password string
	err := s.driver.QueryRowContext(newCtx, getPassword, username).Scan(&password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: no rows found: %w", op, err)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return password, nil
}

func (s *Storage) RegisterUser(ctx context.Context, usr *auth.User) error {
	const op = "storage.postgres.auth.RegisterUser"

	newCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	_, err := s.driver.ExecContext(newCtx, registerUser, usr.Username, usr.Password, usr.Gender, usr.Age)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) IsAdmin(ctx context.Context, username string) (bool, error) {
	const op = "storage.postgres.auth.IsAdmin"

	newCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	var ans bool
	err := s.driver.QueryRowContext(newCtx, isAdmin, username).Scan(&ans)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return ans, nil
}

func (s *Storage) DeleteUser(ctx context.Context, username string) error {
	const op = "storage.postgres.auth.DeleteUser"

	newCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	if _, err := s.driver.ExecContext(newCtx, deleteUser, username); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
