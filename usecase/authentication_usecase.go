package usecase

import (
	"errors"
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
	"github.com/faizalnurrozi/go-crud/pkg/hashing"
)

type AuthenticationUseCase struct {
	*Contract
}

func NewAuthenticationUseCase(contract *Contract) interfaces.IAuthenticationUseCase {
	return &AuthenticationUseCase{Contract: contract}
}

//Function to login and generate jwt token
func (uc AuthenticationUseCase) Login(req *request.LoginRequest) (res view_models.LoginVm, err error) {
	credentialUc := NewCredentialUseCase(uc.Contract)

	//get user data by email
	credential, err := credentialUc.ReadBy("c.email", req.Email, "=")
	if err != nil {
		fmt.Println(err)
		return res, errors.New("CredentialDoNotMatch")
	}

	//validating password
	isValid := hashing.CheckHashString(req.Password, credential.Password)
	if !isValid {
		fmt.Println(err)
		return res, errors.New("CredentialDoNotMatch")
	}

	//generate jwt payload and encrypted with jwe
	payload := map[string]interface{}{
		"id": credential.ID,
	}
	jwePayload, err := uc.JweCredential.GenerateJwePayload(payload)
	if err != nil {
		fmt.Println(err)
		return res, errors.New("CredentialDoNotMatch")
	}

	//generate jwt token
	res, err = uc.GenerateJWT(req.Email, jwePayload)
	if err != nil {
		fmt.Println(err)
		return res, errors.New("CredentialDoNotMatch")
	}

	return res, nil
}

func (AuthenticationUseCase) GoogleOauth(token string) (res view_models.LoginVm, err error) {
	panic("implement me")
}

func (AuthenticationUseCase) FacebookOauth(token string) (res view_models.LoginVm, err error) {
	panic("implement me")
}

func (AuthenticationUseCase) Logout() {
	panic("implement me")
}

func (AuthenticationUseCase) ForgotPassword() (err error) {
	panic("implement me")
}

//
func (uc AuthenticationUseCase) GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error) {
	res.Token, res.TokenExpiredAt, err = uc.JwtCredential.GetToken(issuer, payload)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.RefreshToken, res.RefreshTokenExpiredAt, err = uc.JwtCredential.GetRefreshToken(issuer, payload)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

//
func (uc AuthenticationUseCase) GetCurrentUser() (res view_models.CurrentUserVm, err error) {
	factoryCurrentUser := view_models.NewFactoryCurrentUser(res)

	adminUc := NewAdminUseCase(uc.Contract)
	admin, err := adminUc.ReadBy("a.credential_id", uc.UserID, "=")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res = factoryCurrentUser.AdminCurrentUser(admin, uc.UserID)

	return res, nil
}
