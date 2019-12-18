package inmemory

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/trendyol/teg/pkg/storage"
)

type inMemoryStorage struct {
	featureToggles atomic.Value
	sync.Mutex
}

var _ storage.ReaderWriter = (*inMemoryStorage)(nil)

func New(setters ...option) *inMemoryStorage {
	args := &options{
		DefaultFeatureToggles: make(storage.FeatureToggles),
	}

	for _, setter := range setters {
		setter(args)
	}

	init := inMemoryStorage{}
	init.featureToggles.Store(args.DefaultFeatureToggles)

	return &init
}

func (i *inMemoryStorage) Get(name string) (*storage.FeatureToggle, error) {
	val, ok := i.read(name)
	if !ok {
		return nil, fmt.Errorf("cloud not found %s key", name)
	}

	return &val, nil
}

func (i *inMemoryStorage) GetAll() (storage.FeatureToggles, error) {
	return i.getStorage(), nil
}

func (i *inMemoryStorage) Set(name string, toggle storage.FeatureToggle) error {
	i.insert(name, toggle)

	return nil
}

func (i *inMemoryStorage) getStorage() storage.FeatureToggles {
	return i.featureToggles.Load().(storage.FeatureToggles)
}

func (i *inMemoryStorage) read(name string) (val storage.FeatureToggle, ok bool) {
	featureToggles := i.getStorage()

	val, ok = featureToggles[name]
	return
}

func (i *inMemoryStorage) insert(name string, val storage.FeatureToggle) {
	i.Lock()
	defer i.Unlock()

	featureToggles := i.getStorage()

	newFeatureToggles := make(storage.FeatureToggles)
	for k, v := range featureToggles {
		newFeatureToggles[k] = v
	}
	newFeatureToggles[name] = val

	i.featureToggles.Store(newFeatureToggles)
}
