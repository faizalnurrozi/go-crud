package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type ISalesDetailRepository interface {
	Add(model models.SalesDetail, tx *sql.Tx) (res string, err error)
	DeleteBy(saleID string, tx *sql.Tx) (res int64, err error)
}
