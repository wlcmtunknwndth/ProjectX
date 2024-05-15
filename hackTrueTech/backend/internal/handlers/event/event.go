package event

import (
	"encoding/json"
	"fmt"
	"github.com/wlcmtunknwndth/hackBPA/internal/auth"
	"github.com/wlcmtunknwndth/hackBPA/internal/broker/nats"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/corsSkip"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/httpResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type Cache interface {
	CacheOrder(event storage.Event)
	GetOrder(uuid string) (*storage.Event, bool)
}

type EventsHandler struct {
	Broker *nats.Nats
	Cache  Cache
}

const (
	StatusNotEnoughPermissions = "Not enough permissions"
	StatusUnauthorized         = "Unauthorized"
	StatusBadRequest           = "Bad request"
	StatusEventCreated         = "Event created"
	StatusInternalServerError  = "Internal server error"
	StatusDeleted              = "Event deleted"
	StatusPatched              = "Event patched"
)

func (e *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.event.CreateEvent"

	if !checkAdminRights(w, r) {
		return
	}

	corsSkip.EnableCors(w, r)
	body := r.Body
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			slog.Error("couldn't close request body", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		}
	}(body)

	data, err := io.ReadAll(body)
	if err != nil {
		slog.Error("couldn't read body", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusBadRequest, StatusBadRequest)
		return
	}
	var event storage.Event
	if err = json.Unmarshal(data, &event); err != nil {
		slog.Error("couldn't decode body", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusBadRequest, StatusBadRequest)
		return
	}

	id, err := e.Broker.AskSave(&event)
	if err != nil {
		slog.Error("couldn't send event to broker", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusInternalServerError, StatusInternalServerError)
		return
	}
	event.Id = id
	e.Cache.CacheOrder(event)

	httpResponse.Write(w, http.StatusCreated, fmt.Sprintf("%s: id: %d", StatusEventCreated, id))
}

func (e *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.event.GetEvent"
	corsSkip.EnableCors(w, r)
	event, found := e.Cache.GetOrder(r.URL.Query().Get("id"))
	if found {
		data, err := json.Marshal(event)
		if err != nil {
			slog.Error("couldn't marshall event from cacher", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		}
		if _, err = w.Write(data); err != nil {
			slog.Error("couldn't send event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			httpResponse.Write(w, http.StatusInternalServerError, StatusInternalServerError)
			return
		}
	}

	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		slog.Error("couldn't get event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusBadRequest, StatusBadRequest)
		return
	}
	data, err := e.Broker.AskEvent(uint(id))
	if err != nil {
		slog.Error("couldn't get event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusInternalServerError, StatusInternalServerError)
		return
	}
	time.Sleep(time.Second)

	//w.WriteHeader(http.StatusOK)
	if _, err = w.Write(data); err != nil {
		slog.Error("couldn't send event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusInternalServerError, StatusInternalServerError)
		return
	}
}

func (e *EventsHandler) PatchEvent(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.event.PatchEvent"

	if !checkAdminRights(w, r) {
		return
	}

	corsSkip.EnableCors(w, r)
	body := r.Body
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			slog.Error("couldn't close request body", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		}
	}(body)

	data, err := io.ReadAll(body)
	if err != nil {
		slog.Error("couldn't read body", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusBadRequest, StatusBadRequest)
		return
	}
	var event storage.Event
	if err = json.Unmarshal(data, &event); err != nil {
		slog.Error("couldn't decode body", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusBadRequest, StatusBadRequest)
		return
	}

	if err = e.Broker.AskPatch(&event); err != nil {
		slog.Error("couldn't publish patch ask", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusInternalServerError, StatusInternalServerError)
		return
	}

	httpResponse.Write(w, http.StatusOK, StatusPatched)
}

func (e *EventsHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.event.DeleteEvent"
	if !checkAdminRights(w, r) {
		return
	}

	corsSkip.EnableCors(w, r)

	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		slog.Error("couldn't parse query", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusBadRequest, StatusBadRequest)
		return
	}

	if err = e.Broker.AskDelete(uint(id)); err != nil {
		slog.Error("couldn't delete event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusInternalServerError, StatusInternalServerError)
		return
	}

	httpResponse.Write(w, http.StatusOK, StatusDeleted)
}

func checkAdminRights(w http.ResponseWriter, r *http.Request) bool {
	const op = "handlers.event.checkAdminRights"
	if res, err := auth.Access(r); err != nil {
		if !res {
			slog.Error("not authorized", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			httpResponse.Write(w, http.StatusUnauthorized, StatusUnauthorized)
			return false
		}

		slog.Error("couldn't access user", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusUnauthorized, StatusUnauthorized)
		return false
	}

	if res, err := auth.IsAdmin(r); !res || err != nil {
		if !res {
			slog.Error("not enough permissions", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			httpResponse.Write(w, http.StatusForbidden, StatusNotEnoughPermissions)
			return false
		}
		slog.Error("couldn't access user", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		httpResponse.Write(w, http.StatusUnauthorized, StatusUnauthorized)
		return false
	}

	return true
}
