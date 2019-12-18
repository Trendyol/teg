package main

import (
	"context"
	"fmt"
	"os"

	"github.com/trendyol/teg"
	"github.com/trendyol/teg/pkg/storage/env"
)

func main() {
	os.Setenv("TEG_ENABLE", "true")
	defer os.Unsetenv("TEG_ENABLE")

	featureToggle := teg.New(context.Background(), teg.WithStorage(
		env.New(),
	))

	fmt.Println(featureToggle.Get("Enable"))
}
