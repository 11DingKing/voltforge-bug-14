package power

import (
	"context"
	"sync"
	"testing"
)

func TestVoltForge14(t *testing.T) {
	ledger := NewBudgetLease(100)
	gate := make(chan struct{})
	ready := make(chan struct{}, 2)
	ledger.BeforeCommit = func() { ready <- struct{}{}; <-gate }
	var wg sync.WaitGroup
	errs := make(chan error, 2)
	for n := 0; n < 2; n++ {
		wg.Add(1)
		go func() { defer wg.Done(); errs <- ledger.Reserve(context.Background(), 60) }()
	}
	<-ready
	<-ready
	close(gate)
	wg.Wait()
	close(errs)
	success := 0
	for err := range errs {
		if err == nil {
			success++
		}
	}
	if success != 1 || ledger.Used() != 60 {
		t.Fatalf("success=%d used=%d", success, ledger.Used())
	}
}
