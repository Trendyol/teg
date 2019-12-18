package main

import (
	"context"
	"fmt"

	"github.com/trendyol/teg"
	"github.com/trendyol/teg/pkg/storage"
	"github.com/trendyol/teg/pkg/storage/inmemory"
)

func main() {
	featureToggle := teg.New(context.Background(), teg.WithStorage(
		inmemory.New(inmemory.WithDefaultToggles(storage.FeatureToggles{
			"Enable": storage.FeatureToggle{
				Value: true,
			},
		})),
	))

	fmt.Println(featureToggle.Get("Enable"))
}
