package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type ICustomerRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.Customer, err error)

	ReadBy(column, value, operator string) (res models.Customer, err error)

	Add(model models.Customer, tx *sql.Tx) (res string, err error)

	Edit(model models.Customer, tx *sql.Tx) (res int64, err error)

	DeleteBy(column, value, operator string, model models.Customer, tx *sql.Tx) (res int64, err error)

	Count(search string) (res int, err error)
}
