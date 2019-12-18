package teg

import (
	"context"
	"reflect"
	"testing"

	"github.com/trendyol/teg/pkg/filter"
	"github.com/trendyol/teg/pkg/storage"
)

type emptyStorageReaderWriter struct{}

func (e *emptyStorageReaderWriter) Get(_ string) (*storage.FeatureToggle, error) { return nil, nil }
func (e *emptyStorageReaderWriter) GetAll() (storage.FeatureToggles, error)      { return nil, nil }
func (e *emptyStorageReaderWriter) Set(_ string, _ storage.FeatureToggle) error  { return nil }

func TestWithStorage(t *testing.T) {
	opts := options{}

	emptyStorageReaderWriterPointer := &emptyStorageReaderWriter{}
	WithStorage(emptyStorageReaderWriterPointer)(&opts)

	if !reflect.DeepEqual(opts.storageReader, emptyStorageReaderWriterPointer) {
		t.Errorf("expected storage pointer %p but got %p", opts.storageReader, emptyStorageReaderWriterPointer)
	}
}

func TestWithFilters(t *testing.T) {
	opts := options{}

	filters := []filter.Func{func(_ storage.FeatureToggle) bool { return false }}
	WithFilters(filters...)(&opts)

	if !reflect.DeepEqual(opts.Filters, filters) {
		t.Errorf("expected filters pointer %p but got %p", opts.Filters, filters)
	}
}

func TestWithSyncTrigger(t *testing.T) {
	opts := options{}

	syncTriggerFunc := func(_ context.Context, _ chan<- struct{}) error { return nil }
	WithSyncTrigger(syncTriggerFunc)(&opts)

	if opts.SyncTrigger == nil {
		t.Errorf("expected sync trigger func assigned to pointer but got nil")
	}
}
