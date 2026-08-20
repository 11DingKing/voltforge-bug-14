package power

func (l *BudgetLeaseLedger) Used() int      { return l.used }
func (l *BudgetLeaseLedger) Available() int { return l.capacity - l.used }
