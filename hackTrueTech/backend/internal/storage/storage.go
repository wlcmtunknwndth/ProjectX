package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

type Storage struct {
	db     *sql.DB
	broker nats.Conn
}

type Event struct {
	Id           uint      `json:"id,omitempty"`
	Price        uint      `json:"price"`
	Restrictions uint      `json:"restrictions"`
	Date         time.Time `json:"date"`
	Feature      string    `json:"feature,omitempty"`
	City         string    `json:"city"`
	Address      string    `json:"address"`
	Name         string    `json:"name"`
	ImgPath      string    `json:"img_path"`
	Description  string    `json:"description"`
}

func EventToJSON(event *Event) ([]byte, error) {
	const op = "storage.EventToJSON"
	data, err := json.Marshal(*event)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return data, nil
}

//type Features struct {
//	Disability bool `json:"disability"`
//	Deaf       bool `json:"deaf"`
//	Blind      bool `json:"blind"`
//	Neural     bool `json:"neural"`
//}
