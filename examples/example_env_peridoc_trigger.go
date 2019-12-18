package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/trendyol/teg"
	"github.com/trendyol/teg/pkg/storage/env"
	"github.com/trendyol/teg/pkg/sync"
)

func main() {
	os.Setenv("TEG_ENABLE", "true")
	defer os.Unsetenv("TEG_ENABLE")

	featureToggle := teg.New(context.Background(),
		teg.WithStorage(env.New()),
		teg.WithSyncTrigger(sync.PeriodicTrigger(5*time.Minute)),
	)

	fmt.Println(featureToggle.Get("Enable"))
}
