package request

import "time"

type TransactionStockRequest struct {
	TransactionID string    `json:"transaction_id" validate:"required"`
	Date          time.Time `json:"date" validate:"required"`
	ProductID     string    `json:"product_id" validate:"required"`
	MerchantID    string    `json:"merchant_id" validate:"required"`
	StockIn       int       `json:"stock_in"`
	StockOut      int       `json:"stock_out"`
}
