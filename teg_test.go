package teg

import (
	"context"
	"errors"
	"testing"

	"github.com/trendyol/teg/pkg/storage"
	"github.com/trendyol/teg/pkg/storage/inmemory"
)

func TestTeg(t *testing.T) {
	tg := New(context.Background())
	if tg.Get("WrongName") {
		t.Error("expected false but got true")
	}

	tg = New(context.Background(), WithStorage(
		inmemory.New(inmemory.WithDefaultToggles(storage.FeatureToggles{
			"Name": storage.FeatureToggle{
				Value: true,
			},
		})),
	))

	if !tg.Get("Name") {
		t.Error("expected true but got false")
	}
}

type errorStorageReaderWriter struct{}

func (e *errorStorageReaderWriter) Get(_ string) (*storage.FeatureToggle, error) { return nil, nil }
func (e *errorStorageReaderWriter) GetAll() (storage.FeatureToggles, error) {
	return nil, errors.New("system error")
}
func (e *errorStorageReaderWriter) Set(_ string, _ storage.FeatureToggle) error { return nil }

func TestTegWithSyncError(t *testing.T) {
	tg := New(context.Background(), WithStorage(&errorStorageReaderWriter{}))
	if tg.Get("WrongName") {
		t.Error("expected false but got true")
	}
}
