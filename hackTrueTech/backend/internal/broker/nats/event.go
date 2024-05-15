package nats

import (
	"context"
	"encoding/binary"
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
	MustSaveEvent   = "save_event"
	MustDeleteEvent = "del.*"
	AskDeleteEvent  = "del."
	MustSendEvent   = "get.*"
	AskGetEvent     = "get."
	MustPatchEvent  = "patch.*"
	AskPatchEvent   = "patch."
)

func convertUintToString(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func convertUintToByte(num uint) []byte {
	var data = []byte{byte(0)}
	if num == 0 {
		return data
	}
	binary.BigEndian.PutUint64(data, uint64(num))
	return data
}

func convertStrToUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func (n *Nats) EventSender(ctx context.Context) (*nats.Subscription, error) {
	const op = "broker.nats.event.PublishEvent"
	sub, err := n.b.Subscribe(MustSendEvent, func(msg *nats.Msg) {
		id, err := convertStrToUint(msg.Subject[4:])
		if err != nil {
			slog.Error("couldn't convert str to uint", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}

		event, err := n.db.GetEvent(ctx, uint(id))
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

func (n *Nats) AskEvent(id uint) ([]byte, error) {
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

func (n *Nats) AskSave(event *storage.Event) (uint, error) {
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
	return uint(id), nil
}

func (n *Nats) EventDeleter(ctx context.Context) (*nats.Subscription, error) {
	const op = "broker.nats.event.Deleter"
	sub, err := n.b.Subscribe(MustDeleteEvent, func(msg *nats.Msg) {
		id, err := strconv.ParseUint(msg.Subject[4:], 10, 64)
		if err != nil {
			slog.Error("couldn't parse id", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
		if err = n.db.DeleteEvent(ctx, uint(id)); err != nil {
			slog.Error("couldn't delete event", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (n *Nats) AskDelete(id uint) error {
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
