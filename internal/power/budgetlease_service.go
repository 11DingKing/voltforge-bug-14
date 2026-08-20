package power

import (
	"context"
	"errors"
	"sync"
)

var ErrBudgetLeaseCapacity = errors.New("budgetlease capacity exceeded")

type BudgetLeaseLedger struct {
	mu             sync.Mutex
	capacity, used int
	BeforeCommit   func()
}

func NewBudgetLease(capacity int) *BudgetLeaseLedger { return &BudgetLeaseLedger{capacity: capacity} }

func (l *BudgetLeaseLedger) Reserve(ctx context.Context, amount int) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	l.mu.Lock()
	if l.used+amount > l.capacity {
		l.mu.Unlock()
		return ErrBudgetLeaseCapacity
	}
	hook := l.BeforeCommit
	l.mu.Unlock()
	if hook != nil {
		hook()
	}
	l.mu.Lock()
	l.used += amount
	l.mu.Unlock()
	return nil
}
