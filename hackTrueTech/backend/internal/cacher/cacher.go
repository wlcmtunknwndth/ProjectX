package cacher

import (
	"github.com/patrickmn/go-cache"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"log/slog"
	"strconv"
	"time"
)

const cacheExpiration = 2 * time.Minute

// cached -- is the map with saved uuids in current run, so it is easier to back up
//var cached = make(map[uint]struct{})

type Storage interface {
	RestoreCache() ([]storage.Event, error)
	SaveCache(id uint) error
	DeleteCache(id uint) error
	IsAlreadyCached(id uint) bool
}

type Cacher struct {
	handler *cache.Cache
	db      Storage
}

// New -- creates new instance of Cacher with Storage interface and cacher.Cache vars. expTime -- is the standard expiration time of cached item.
// purgeTime -- is the time the cacher cleans up itself
func New(db Storage, expTime time.Duration, purgeTime time.Duration) *Cacher {
	return &Cacher{
		handler: cache.New(expTime, purgeTime),
		db:      db,
	}
}

func formatUint(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func formatString(numStr string) uint {
	num, _ := strconv.ParseUint(numStr, 10, 64)
	return uint(num)
}

// CacheOrder -- caches the order given as an arg and maps order's uuid to cacher map.
func (c *Cacher) CacheOrder(event storage.Event) {
	c.handler.OnEvicted(c.onEvicted)
	slog.Info("successfully cached event: ", slog.Attr{Key: event.Name})
	c.handler.Set(formatUint(event.Id), event, cacheExpiration)
	//cached[event.Id] = struct{}{}
}

// onEvicted -- is a custom func, handling cached item after expiration. It deletes item from cacher map and deletes uuid from storage Cache backup.
func (c *Cacher) onEvicted(id string, data interface{}) {
	//delete(cached, formatString(id))
	err := c.db.DeleteCache(formatString(id))
	if err != nil {
		slog.Error("couldn't delete order from cacher")
	}
}

// GetOrder -- gets order from cacher if found
func (c *Cacher) GetOrder(id string) (*storage.Event, bool) {
	data, found := c.handler.Get(id)
	if found {
		order := data.(storage.Event)
		slog.Info("successfully found event", slog.Attr{Key: order.Name})
		return &order, true
	}
	return nil, false
}

// Restore -- restores cached item from backup copy in storage. Must be used at the start of ur application.
func (c *Cacher) Restore() error {
	orders, err := c.db.RestoreCache()
	//fmt.Println(orders)
	if err != nil {
		slog.Error("couldn't restore cacher: ", err)
		return err
	}

	for i := range orders {
		c.CacheOrder(orders[i])
	}
	return nil
}

// SaveCache -- backups cacher to the storage
func (c *Cacher) SaveCache() error {
	var err error
	for key := range c.handler.Items() {
		if c.db.IsAlreadyCached(formatString(key)) {
			continue
		}
		err = c.db.SaveCache(formatString(key))
		if err != nil {
			slog.Error("couldn't save uuid to cacher zone: ", key, err)
			continue
		}
	}
	return nil
}
