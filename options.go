package teg

import (
	"github.com/trendyol/teg/pkg/filter"
	"github.com/trendyol/teg/pkg/storage"
	"github.com/trendyol/teg/pkg/sync"
)

type options struct {
	storageReader storage.ReaderWriter
	Filters       []filter.Func
	SyncTrigger   sync.TriggerFunc
}

type option func(*options)

func WithStorage(val storage.ReaderWriter) option {
	return func(args *options) {
		args.storageReader = val
	}
}

func WithFilters(val ...filter.Func) option {
	return func(args *options) {
		args.Filters = val
	}
}

func WithSyncTrigger(val sync.TriggerFunc) option {
	return func(args *options) {
		args.SyncTrigger = val
	}
}
