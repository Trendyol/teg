package sync

import (
	"context"
	"testing"
	"time"
)

func TestPeriodicSyncTrigger(t *testing.T) {
	syncChan := make(chan struct{}, 1)

	// Correct Test Case
	ctx, cancel := context.WithTimeout(context.Background(), 110*time.Millisecond)
	go PeriodicTrigger(100*time.Millisecond)(ctx, syncChan)

	for {
		select {
		case <-syncChan:
			break
		case <-ctx.Done():
			t.Errorf("couldn't receive any trigger data")
			break
		}
		break
	}

	cancel()

	// Incorrect Test Case
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go PeriodicTrigger(1000*time.Millisecond)(context.Background(), syncChan)

	for {
		select {
		case <-syncChan:
			t.Errorf("received trigger data")
			break
		case <-ctx.Done():
			break
		}
		break
	}

}
