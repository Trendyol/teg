package main

import (
	"context"
	"fmt"

	"github.com/trendyol/teg"
	"github.com/trendyol/teg/pkg/storage/git"
)

func main() {
	featureToggle := teg.New(context.Background(), teg.WithStorage(
		git.New("https://github.com/trendyol/teg", "examples/feature_toogles.yaml"),
	))

	fmt.Println(featureToggle.Get("Enable"))
}
