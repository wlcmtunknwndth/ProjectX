package postgres

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/compareStrings"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"slices"
	"strings"
	"time"
)

const (
	imageFolder = "/data/events"
	featureSep  = "/"
)

func (s *Storage) GetEvent(ctx context.Context, id uint) (*storage.Event, error) {
	const op = "storage.postgres.events.GetEvent"

	var event storage.Event

	err := s.driver.QueryRowContext(ctx, getEvent, id).Scan(
		&event.Id, &event.Price, &event.Restrictions, &event.Date,
		&event.Feature, &event.City, &event.Address, &event.Name,
		&event.ImgPath, &event.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &event, nil
}

func (s *Storage) getId(name string, date time.Time) (uint, error) {
	const op = "storage.postgres.events.getId"

	var id uint
	if err := s.driver.QueryRow(getId, name, date).Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) CreateEvent(ctx context.Context, event *storage.Event) (uint, error) {
	const op = "storage.postgres.events.CreateEvent"

	if len(event.Feature) != 0 {
		features := strings.Split(event.Feature, featureSep)
		if len(features) != 0 {
			slices.SortFunc(features, compareStrings.CmpStr)
			event.Feature = strings.Join(features, featureSep)
		}
	}

	_, err := s.driver.ExecContext(ctx, createEvent, &event.Price,
		&event.Restrictions, &event.Date, &event.Feature, &event.City,
		&event.Address, &event.Name, imageFolder, &event.Description,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id uint
	if id, err = s.getId(event.Name, event.Date); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) DeleteEvent(ctx context.Context, id uint) error {
	const op = "storage.postgres.events.DeleteEvent"

	if _, err := s.driver.ExecContext(ctx, deleteEvent, id); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) PatchEvent(ctx context.Context, event *storage.Event) error {
	const op = "storage.postgres.events.PatchEvent"

	_, err := s.driver.ExecContext(ctx, patchEvent, &event.Price,
		&event.Restrictions, &event.Date, &event.Feature, &event.City,
		&event.Address, &event.Name, &event.Description, &event.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

//func (s *Storage) GetEventsByFeature(ctx context.Context, date time.Time, feature string) ([]storage.Event, error) {
//	if len(feature) != 0 {
//		features := strings.Split(feature, featureSep)
//		if len(features) != 0 {
//			slices.SortFunc(features, compareStrings.CmpStr)
//			feature = strings.Join(features, featureSep)
//		}
//	}
//	rows, err := s.driver.QueryRowContext()
//
//}
