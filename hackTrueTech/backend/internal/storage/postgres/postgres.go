package postgres

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
)

type Storage struct {
	driver *sql.DB
}

type Index struct {
	Id        uint64
	EventId   uint64
	FeatureId pq.Int64Array
}

func New(config *config.Database) (*Storage, error) {
	const op = "storage.postgres.New"

	connStr := fmt.Sprintf("postgres://%s:%s@postgres:%s/%s?sslmode=%s", config.DbUser, config.DbPass, config.Port, config.DbName, config.SslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{driver: db}, nil
}

func (s *Storage) Close() error {
	return s.driver.Close()
}

func (s *Storage) Ping() error {
	return s.driver.Ping()
}
