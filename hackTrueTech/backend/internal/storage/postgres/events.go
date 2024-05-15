package postgres

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"slices"
)

var featuresToId = map[string]int{
	"blind":      1,
	"deaf":       2,
	"disability": 3,
	"neuro":      4,
}

var idToFeature = map[uint64]string{
	1: "blind",
	2: "deaf",
	3: "disability",
	4: "neuro",
}

func (s *Storage) GetEvent(ctx context.Context, id uint64) (*storage.Event, error) {
	const op = "storage.postgres.events.GetEvent"

	var index storage.Index
	err := s.driver.QueryRowContext(ctx, getIndex, &id).Scan(&index.EventId, pq.Array(&index.FeatureId))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var event storage.Event
	err = s.driver.QueryRowContext(ctx, getEvent, index.EventId).Scan(
		&event.Id, &event.Price, &event.Restrictions, &event.Date,
		&event.City, &event.Address, &event.Name,
		&event.ImgPath, &event.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for _, val := range index.FeatureId {
		ftr, ok := idToFeature[uint64(val.Int64)]
		if !ok {
			continue
		}
		event.Feature = append(event.Feature, ftr)
	}

	return &event, nil
}

func cmpStrings(a, b string) int {
	if len(a) == 0 && len(b) == 0 {
		return 0
	} else if len(a) == 0 {
		return -1
	} else if len(b) == 0 {
		return 1
	}
	for i := 0; i < len(a) || i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}
	return 0
}

func (s *Storage) CreateEvent(ctx context.Context, event *storage.Event) (uint64, error) {
	const op = "storage.postgres.events.CreateEvent"

	var id uint64
	err := s.driver.QueryRowContext(ctx, createEvent, &event.Price,
		&event.Restrictions, &event.Date, &event.City,
		&event.Address, &event.Name, &event.ImgPath, &event.Description,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var features = make([]int, 0, 2)
	if event.Feature != nil {
		slices.SortFunc(event.Feature, cmpStrings)
		for _, val := range event.Feature {
			featureId, ok := featuresToId[val]
			if !ok {
				continue
			}
			features = append(features, featureId)
		}
	}

	var indId uint64
	if err = s.driver.QueryRowContext(ctx, createIndex, &id, pq.Array(features)).Scan(&indId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return indId, nil
}

func (s *Storage) DeleteEvent(ctx context.Context, id uint64) error {
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
