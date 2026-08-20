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
	if hook := l.BeforeCommit; hook != nil {
		hook()
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if amount <= 0 || l.used+amount > l.capacity {
		return ErrBudgetLeaseCapacity
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	l.used += amount
	return nil
}
