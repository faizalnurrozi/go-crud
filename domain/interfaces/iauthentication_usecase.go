package interfaces

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
)

type IAuthenticationUseCase interface {
	Login(req *request.LoginRequest) (res view_models.LoginVm, err error)

	Logout()

	GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error)

	GetCurrentUser() (res view_models.CurrentUserVm, err error)
}
