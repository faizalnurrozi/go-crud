package view_models

import (
	"github.com/faizalnurrozi/go-crud/domain/models"
	"time"
)

type TransactionStockVm struct {
	TransactionID string    `json:"transaction_id"`
	Date          time.Time `json:"date"`
	ProductID     string    `json:"product_id"`
	MerchantID    string    `json:"merchant_id"`
	StockIn       int       `json:"stock_in"`
	StockOut      int       `json:"stock_out"`
}

func (vm TransactionStockVm) Build(model models.TransactionStock) TransactionStockVm {
	return TransactionStockVm{
		TransactionID: model.TransactionID,
		Date:          model.Date,
		ProductID:     model.ProductID,
		MerchantID:    model.MerchantID,
		StockIn:       model.StockIn,
		StockOut:      model.StockOut,
	}
}
