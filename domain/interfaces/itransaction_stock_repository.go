package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type ITransactionStockRepository interface {
	Add(model models.TransactionStock, tx *sql.Tx) (res string, err error)
	Diff(merchantId, outletId, productId string) (res int, err error)
}
