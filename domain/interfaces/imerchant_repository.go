package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type IMerchantRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.Merchant, err error)

	ReadBy(column, value, operator string) (res models.Merchant, err error)

	Add(model models.Merchant, tx *sql.Tx) (res string, err error)

	Edit(model models.Merchant, tx *sql.Tx) (res int64, err error)

	DeleteBy(column, value, operator string, model models.Merchant, tx *sql.Tx) (res int64, err error)

	Count(search string) (res int, err error)
}
