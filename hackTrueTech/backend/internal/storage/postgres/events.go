package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/compareStrings"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"log/slog"
	"slices"
	"sync"
)

var featuresToId = map[string]int{
	"blind":      1,
	"deaf":       2,
	"disability": 3,
	"neuro":      4,
}

var idToFeature = map[int64]string{
	1: "blind",
	2: "deaf",
	3: "disability",
	4: "neuro",
}

func (s *Storage) GetEvent(ctx context.Context, id uint64) (*storage.Event, error) {
	const op = "storage.postgres.events.GetEvent"

	var index Index
	err := s.driver.QueryRowContext(ctx, getIndex, &id).Scan(&index.EventId, &index.FeatureId)
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

	features := make([]string, 0, len(index.FeatureId))
	for i, _ := range index.FeatureId {
		//var idd int64
		//err = index.FeatureId
		//if err != nil {
		//	slog.Error("couldn't scan id", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		//	continue
		//}
		ftr, ok := idToFeature[index.FeatureId[i]]
		if !ok {
			continue
		}
		features = append(features, ftr)
	}
	event.Feature = features

	return &event, nil
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
		slices.SortFunc(event.Feature, compareStrings.CmpStr)
		for _, val := range event.Feature {
			featureId, ok := featuresToId[val]
			if ok {
				features = append(features, featureId)
			}
		}
	}

	var indId uint64
	if err = s.driver.QueryRowContext(ctx, createIndex, &id, pq.Array(features)).Scan(&indId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return indId, nil
}

func (s *Storage) GetEventsByFeature(ctx context.Context, features []string) ([]storage.Event, error) {
	const op = "storage.postgres.events.GetEventsByFeature"

	var rows *sql.Rows
	if features != nil {
		var ids = make([]int, 0, len(features))
		for i, _ := range features {
			id, ok := featuresToId[features[i]]
			if ok {
				ids = append(ids, id)
			}
		}
		if len(ids) == 0 {
			ids = []int{1, 2, 3, 4}
		}
		var err error
		rows, err = s.driver.QueryContext(ctx, getIndexesByFeature, pq.Array(ids))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

	} else {
		var err error
		rows, err = s.driver.QueryContext(ctx, getAllIndexes)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)

		}
	}

	var events []storage.Event
	var wg sync.WaitGroup
	for rows.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var index Index
			err := rows.Scan(&index.Id, &index.EventId, &index.FeatureId)
			if err != nil {
				slog.Error("couldn't scan index row", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
				return
			}

			var event storage.Event
			if err := s.driver.QueryRow(getEvent, index.Id).Scan(&event.Id, &event.Price,
				&event.Restrictions, &event.Date, &event.City, &event.Address, &event.Name,
				&event.ImgPath, &event.Description,
			); err != nil {
				slog.Error("couldn't get event by feature", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
				return
			}
			events = append(events, event)
		}()
		wg.Wait()
	}

	return events, nil
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
