package models

import (
	"time"
)

type TransactionStock struct {
	ID            string
	TransactionID string
	Date          time.Time
	ProductID     string
	MerchantID    string
	StockIn       int
	StockOut      int
}
