package sync

import (
	"context"
)

func Start(ctx context.Context, triggerFunc TriggerFunc, syncFunc func()) {
	syncFunc()

	syncChan := make(chan struct{})
	go triggerFunc(ctx, syncChan)
	go listener(ctx, syncChan, syncFunc)
}

func listener(ctx context.Context, syncListener <-chan struct{}, callback func()) {
	for {
		select {
		case <-syncListener:
			callback()
		case <-ctx.Done():
			return
		}
	}
}
