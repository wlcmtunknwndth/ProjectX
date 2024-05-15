package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
	"sync"
)

type Storage struct {
	driver *sql.DB
	mtx    sync.RWMutex
}

func New(config *config.Database) (*Storage, error) {
	const op = "storage.postgres.New"

	connStr := fmt.Sprintf("postgres://%s:%s@postgres:%s/%s?sslmode=%s", config.DbUser, config.DbPass, config.Port, config.DbName, config.SslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
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

//func (s *Storage) Lock() {
//	s.mtx.Lock()
//}
//
//func (s *Storage) Unlock() {
//	s.mtx.Unlock()
//}
//
//func (s *Storage) RLock() {
//	s.mtx.RLock()
//}
//
//func (s *Storage) RUnlock() {
//	s.mtx.RUnlock()
//}
