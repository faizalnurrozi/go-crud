package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type ISupplierRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.Supplier, err error)

	ReadBy(column, value, operator string) (res models.Supplier, err error)

	Add(model models.Supplier, tx *sql.Tx) (res string, err error)

	Edit(model models.Supplier, tx *sql.Tx) (res int64, err error)

	DeleteBy(column, value, operator string, model models.Supplier, tx *sql.Tx) (res int64, err error)

	Count(search string) (res int, err error)
}
