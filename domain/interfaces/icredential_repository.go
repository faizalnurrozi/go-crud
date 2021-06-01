package interfaces

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type ICredentialRepository interface {
	ReadBy(column, value, operator string) (res models.Credential, err error)

	Add(model models.Credential, tx *sql.Tx) (res string, err error)

	DeleteBy(column, value, operator string, model models.Credential, tx *sql.Tx) (err error)

	CountBy(column, value, operator string) (res int, err error)
}
