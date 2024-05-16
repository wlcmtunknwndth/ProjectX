package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ImageFolder = "./data"
)

type Storage struct {
	db     *sql.DB
	broker nats.Conn
}

type Event struct {
	Id           uint64    `json:"id,omitempty"`
	Price        uint64    `json:"price"`
	Restrictions uint64    `json:"restrictions"`
	Date         time.Time `json:"date"`
	Feature      []string  `json:"feature,omitempty"`
	City         string    `json:"city"`
	Address      string    `json:"address"`
	Name         string    `json:"name"`
	ImgPath      string    `json:"img_path"`
	Description  string    `json:"description"`
}

type Index struct {
	Id        uint64
	EventId   uint64
	FeatureId []sql.NullInt64
}

func EventToJSON(event *Event) ([]byte, error) {
	const op = "storage.EventToJSON"
	data, err := json.Marshal(*event)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return data, nil
}

const (
	MaxSizeForm = 1024 * 1024 * 10
	ImageExt    = ".svg"
)

func ParseFormData(r *http.Request) (*Event, error) {
	const op = "handlers.event.parseFormData"
	err := r.ParseMultipartForm(MaxSizeForm)
	mForm := r.MultipartForm
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var event Event

	if event.Price, err = strconv.ParseUint(mForm.Value["price"][0], 10, 64); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if event.Restrictions, err = strconv.ParseUint(mForm.Value["restrictions"][0], 10, 64); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if event.Date, err = time.Parse(time.RFC3339, mForm.Value["date"][0]); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if event.Feature = mForm.Value["feature"]; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if event.City = mForm.Value["city"][0]; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if event.Address = mForm.Value["address"][0]; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if event.Name = mForm.Value["name"][0]; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if event.Description = mForm.Value["description"][0]; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	year, month, day := event.Date.Date()
	path := fmt.Sprintf("%s/%d%d%d%s", ImageFolder, year, month, day, event.Name)
	if err = os.Mkdir(path, 0777); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, file := range mForm.File["img_path"] {
		img, err := file.Open()
		if err != nil {
			slog.Error("couldn't open sent image", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			continue
		}
		ext := strings.Split(file.Filename, ".")
		localFile, err := os.Create(fmt.Sprintf("%s/%d%s", path, i, ext[len(ext)-1]))
		if err != nil {
			slog.Error("couldn't create image copy", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			continue
		}
		if _, err = io.Copy(localFile, img); err != nil {
			slog.Error("couldn't save image", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			continue
		}
	}

	event.ImgPath = path

	return &event, nil
}
