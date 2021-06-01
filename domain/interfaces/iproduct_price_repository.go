package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type IProductPriceRepository interface {
	ReadBy(outletId, productId string) (res models.ProductPrice, err error)

	Add(model models.ProductPrice, tx *sql.Tx) (res string, err error)

	Edit(model models.ProductPrice, tx *sql.Tx) (res int64, err error)

	DeleteBy(outletId, productId string, tx *sql.Tx) (res int64, err error)
}
