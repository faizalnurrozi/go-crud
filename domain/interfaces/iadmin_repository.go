package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type IAdminRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.Admin, err error)

	ReadBy(column, value, operator string) (res models.Admin, err error)

	Add(model models.Admin, tx *sql.Tx) (res string, err error)

	Edit(model models.Admin, tx *sql.Tx) (res int64, err error)

	DeleteBy(column, value, operator string, model models.Admin, tx *sql.Tx) (res int64, err error)

	Count(search string) (res int, err error)
}
