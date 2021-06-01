package interfaces

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
)

type ITransactionStockUseCase interface {
	Add(req *request.TransactionStockRequest) (res string, err error)
	Diff(merchantId, outletId, productId string) (res int, err error)
}
