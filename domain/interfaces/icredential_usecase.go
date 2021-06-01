package interfaces

import (
	"github.com/faizalnurrozi/go-crud/domain/view_models"
)

type ICredentialUseCase interface {
	ReadBy(column, value, operator string) (res view_models.CredentialVm, err error)

	Add(email, password string) (res string, err error)

	DeleteBy(column, value, operator string) (err error)

	CountBy(column, value, operator string) (res int, err error)

	ValidateCredential(column, value, operator, password string) (res bool, err error)
}
