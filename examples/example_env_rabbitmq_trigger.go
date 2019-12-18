package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/trendyol/teg"
	"github.com/trendyol/teg/pkg/storage/env"
	"github.com/trendyol/teg/pkg/sync"
)

func main() {
	os.Setenv("TEG_ENABLE", "true")
	defer os.Unsetenv("TEG_ENABLE")

	rabbitmqTrigger, err := sync.AmqpTrigger("amqp://guest:guest@localhost:5672/", "teg-update")
	if err != nil {
		log.Fatal(err)
	}

	featureToggle := teg.New(context.Background(),
		teg.WithStorage(env.New()),
		teg.WithSyncTrigger(rabbitmqTrigger),
	)

	fmt.Println(featureToggle.Get("Enable"))
}
