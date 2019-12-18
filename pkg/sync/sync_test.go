package sync

import (
	"context"
	"testing"
	"time"
)

func TestSyncListener(t *testing.T) {
	syncChan := make(chan struct{}, 1)
	syncChan <- struct{}{}

	// Correct Test Case
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	isCalled := false
	go listener(ctx, syncChan, func() {
		isCalled = true
	})

	time.Sleep(110 * time.Millisecond)

	if !isCalled && ctx.Err() != nil {
		t.Errorf("callback function is not called")
	}

	cancel()

	// Incorrect Test Case
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	go listener(ctx, syncChan, func() {})
	cancel()

	if ctx.Err() == nil {
		t.Errorf("expeceted error but got nil")
	}
}

func TestStartSync(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	isCalled := false

	Start(ctx, PeriodicTrigger(1*time.Millisecond), func() {
		isCalled = true
	})

	if !isCalled {
		t.Errorf("callback function is not called")
	}
}
