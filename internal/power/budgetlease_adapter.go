package power

func (l *BudgetLeaseLedger) Used() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.used
}
func (l *BudgetLeaseLedger) Available() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.capacity - l.used
}
