package usecase

import (
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/repositories"
	"time"
)

type TransactionStockUseCase struct {
	*Contract
}

func NewTransactionStockUseCase(ucContract *Contract) interfaces.ITransactionStockUseCase {
	return &TransactionStockUseCase{Contract: ucContract}
}

func (uc TransactionStockUseCase) Add(req *request.TransactionStockRequest) (res string, err error) {
	repository := repositories.NewTransactionStockRepository(uc.DB)
	now := time.Now().UTC()

	model := models.TransactionStock{
		TransactionID: req.TransactionID,
		Date:          now,
		ProductID:     req.ProductID,
		MerchantID:    req.MerchantID,
		StockIn:       req.StockIn,
		StockOut:      req.StockOut,
	}

	res, err = repository.Add(model, uc.TX)
	if err != nil {
		fmt.Print("query-add: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

func (uc TransactionStockUseCase) Diff(merchantId, outletId, productId string) (res int, err error) {
	repository := repositories.NewTransactionStockRepository(uc.DB)
	res, err = repository.Diff(merchantId, outletId, productId)
	if err != nil {
		fmt.Print("query-count: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
