package nats

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"log/slog"
	"strconv"
	"time"
)

const (
	MustSaveEvent          = "save_event"
	MustDeleteEvent        = "del.*"
	AskDeleteEvent         = "del."
	MustSendEvent          = "get.*"
	AskGetEvent            = "get."
	MustPatchEvent         = "patch.*"
	AskPatchEvent          = "patch."
	MustSendFilteredEvents = "filtered_events"
	AskFilteredEvents      = "filtered_events"
)

func convertUintToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func convertStrToUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func (n *Nats) EventSender(ctx context.Context) (*nats.Subscription, error) {
	const op = "broker.nats.event.PublishEvent"
	sub, err := n.b.Subscribe(MustSendEvent, func(msg *nats.Msg) {
		id, err := convertStrToUint(msg.Subject[4:])
		if err != nil {
			slog.Error("couldn't convert str to uint64", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		event, err := n.db.GetEvent(ctx, id)
		if err != nil {
			slog.Error("couldn't get event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		data, err := json.Marshal(event)
		if err != nil {
			slog.Error("couldn't marshall event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		if err = msg.Respond(data); err != nil {
			slog.Error("couldn't send reply", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return sub, nil
}

func (n *Nats) AskEvent(id uint64) ([]byte, error) {
	const op = "broker.nats.event.GetEvent"

	msg, err := n.b.Request(AskGetEvent+convertUintToString(id), nil, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return msg.Data, nil
}

func (n *Nats) EventSaver(ctx context.Context) (*nats.Subscription, error) {
	const op = "broker.nats.event.Saver"
	sub, err := n.b.Subscribe(MustSaveEvent, func(msg *nats.Msg) {
		var event storage.Event
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			slog.Error("couldn't unmarshall event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		id, err := n.db.CreateEvent(ctx, &event)
		if err != nil {
			slog.Error("couldn't create event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		if err = msg.Respond([]byte(convertUintToString(id))); err != nil {
			slog.Error("couldn't publish id", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (n *Nats) AskSave(event *storage.Event) (uint64, error) {
	const op = "broker.nats.Event.AskSave"
	data, err := json.Marshal(*event)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	msg, err := n.b.Request(MustSaveEvent, data, 5*time.Second)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := strconv.ParseUint(string(msg.Data), 10, 64)
	if err != nil {
		slog.Error("couldn't atoi", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (n *Nats) EventDeleter(ctx context.Context) (*nats.Subscription, error) {
	const op = "broker.nats.event.Deleter"
	sub, err := n.b.Subscribe(MustDeleteEvent, func(msg *nats.Msg) {
		id, err := strconv.ParseUint(msg.Subject[4:], 10, 64)
		if err != nil {
			slog.Error("couldn't parse id", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
		if err = n.db.DeleteEvent(ctx, id); err != nil {
			slog.Error("couldn't delete event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (n *Nats) AskDelete(id uint64) error {
	return n.b.Publish(fmt.Sprintf("%s%d", AskDeleteEvent, id), nil)
}

func (n *Nats) EventPatcher(ctx context.Context) (*nats.Subscription, error) {
	const op = "broker.nats.event.EventPatcher"
	sub, err := n.b.Subscribe(MustPatchEvent, func(msg *nats.Msg) {
		var event storage.Event
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			return
		}
		if err := n.db.PatchEvent(ctx, &event); err != nil {
			slog.Error("couldn't patch event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return sub, nil
}

func (n *Nats) AskPatch(event *storage.Event) error {
	const op = "broker.nats.event.AskPatch"
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return n.b.Publish(fmt.Sprintf("%s%d", AskPatchEvent, event.Id), data)
}

func (n *Nats) FilteredEventsSender(ctx context.Context) (*nats.Subscription, error) {
	const op = "broker.nats.event.FilteredEventsSender"

	sub, err := n.b.Subscribe(MustSendFilteredEvents, func(msg *nats.Msg) {
		var features []string
		buf := &bytes.Buffer{}
		buf.Write(msg.Data)
		err := gob.NewDecoder(buf).Decode(&features)
		if err != nil {
			slog.Error("couldn't decode features", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		events, err := n.db.GetEventsByFeature(ctx, features)
		if err != nil {
			slog.Error("couldn't get events by feature", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		data, err := json.Marshal(events)
		if err != nil {
			slog.Error("couldn't marshal events", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		if err = msg.Respond(data); err != nil {
			slog.Error("couldn't respond to sub", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return sub, nil
}

func (n *Nats) AskFilteredEvents(features []string) ([]byte, error) {
	const op = "broker.nats.event.AskFilteredEvents"
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(&features)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	msg, err := n.b.Request(AskFilteredEvents, buf.Bytes(), 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return msg.Data, nil
}
