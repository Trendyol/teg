package teg

import (
	"context"
	"log"
	"time"

	"github.com/trendyol/teg/pkg/filter"
	"github.com/trendyol/teg/pkg/storage"
	"github.com/trendyol/teg/pkg/storage/inmemory"
	"github.com/trendyol/teg/pkg/sync"
)

type Teg struct {
	inMemoryStorage storage.ReaderWriter
	filters         []filter.Func
}

func New(ctx context.Context, setters ...option) *Teg {
	args := &options{
		storageReader: inmemory.New(),
		Filters:       filter.RequiredFilters(),
		SyncTrigger:   sync.PeriodicTrigger(1 * time.Minute),
	}

	for _, setter := range setters {
		setter(args)
	}

	init := Teg{
		inMemoryStorage: inmemory.New(),
		filters:         args.Filters,
	}

	// start sync data from storage to in-memory storage
	sync.Start(ctx, args.SyncTrigger, func() {
		if err := storage.Sync(args.storageReader, init.inMemoryStorage); err != nil {
			log.Printf("[ERR] couldn't sync to in-memory storage from storage reader. err: %s\n", err)
		}
	})

	return &init
}

func (t *Teg) Get(name string) bool {
	toggle, err := t.inMemoryStorage.Get(name)
	if err != nil {
		log.Printf("[ERR] couldn't get (%s) feature toggle. err: %s\n", name, err)
		return false
	}

	toogleCopy := *toggle
	return filter.Filter(toogleCopy, t.filters...)
}
