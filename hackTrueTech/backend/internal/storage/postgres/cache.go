package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"log/slog"
)

func (s *Storage) RestoreCache() ([]storage.Event, error) {
	const op = "storage.postgres.RestoreCache"

	rows, err := s.driver.Query(getCachedIds)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	var ids = make([]uint, 0, 5)
	for rows.Next() {
		var id uint
		if err = rows.Scan(&id); err != nil {
			slog.Error("couldn't scan id from cacher table", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			continue
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("%s: %s", op, "no ids cached")
	}

	var events = make([]storage.Event, 0, 3)
	for _, id := range ids {
		var event storage.Event
		if err = s.driver.QueryRow(getEvent, id).Scan(&event.Id, &event.Price, &event.Restrictions, &event.Date,
			&event.Feature, &event.City, &event.Address, &event.Name,
			&event.ImgPath, &event.Description); err != nil {
			slog.Error("couldn't query row id", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			continue
		}
		events = append(events, event)
	}

	return events, nil
}

func (s *Storage) SaveCache(id uint) error {
	const op = "storage.postgres.cacher.SaveCache"
	_, err := s.driver.Exec(saveCache, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteCache(id uint) error {
	const op = "storage.postgres.cacher.DeleteCache"
	_, err := s.driver.Exec(deleteCache, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) IsAlreadyCached(id uint) bool {

	row := s.driver.QueryRow(isAlreadyCached, id)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return true
	}
	return false
}
