package inmemory

import (
	"github.com/trendyol/teg/pkg/storage"
)

type options struct {
	DefaultFeatureToggles storage.FeatureToggles
}

type option func(*options)

func WithDefaultToggles(val storage.FeatureToggles) option {
	return func(args *options) {
		args.DefaultFeatureToggles = val
	}
}
