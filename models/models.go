package models

import "time"

type Expense struct {
	ID          int64
	HostID      int64
	Currency    string
	ExpenseDate time.Time
	Amount      float64
	Customer    string
}
