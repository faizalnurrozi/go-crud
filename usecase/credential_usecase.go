package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
	"github.com/faizalnurrozi/go-crud/pkg/hashing"
	"github.com/faizalnurrozi/go-crud/pkg/messages"
	"github.com/faizalnurrozi/go-crud/repositories"
	"time"
)

type CredentialUseCase struct {
	*Contract
}

func NewCredentialUseCase(UcContract *Contract) interfaces.ICredentialUseCase {
	return &CredentialUseCase{Contract: UcContract}
}

//Function to select credential table with specific where filter
func (uc CredentialUseCase) ReadBy(column, value, operator string) (res view_models.CredentialVm, err error) {
	repository := repositories.NewCredentialRepository(uc.DB)

	credential, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	res = view_models.NewCredentialVm().Build(&credential)

	return res, err
}

//add credential to db
func (uc CredentialUseCase) Add(email, password string) (res string, err error) {
	repository := repositories.NewCredentialRepository(uc.DB)
	now := time.Now().UTC()

	//count credential by email to prevent duplication
	count, err := uc.CountBy("email", email, "=")
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	if count > 0 {
		fmt.Println(err)
		return res, errors.New(messages.DataAlreadyExist)
	}

	//hashing password and add to table credential
	hashed, _ := hashing.HashAndSalt(password)
	model := models.Credential{
		Email:     email,
		Password:  hashed,
		CreatedAt: now,
		UpdatedAt: now,
	}
	res, err = repository.Add(model, uc.TX)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

//delete by specific column
func (uc CredentialUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewCredentialRepository(uc.DB)
	now := time.Now().UTC()

	//count credential by email to prevent error sql no rows result set
	count, err := uc.CountBy(column, value, "=")
	if err != nil {
		fmt.Println(err)
		return err
	}

	if count > 0 {
		model := models.Credential{
			UpdatedAt: now,
			DeletedAt: sql.NullTime{Time: now},
		}
		err = repository.DeleteBy(column, value, operator, model, uc.TX)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

//count by specific column
func (uc CredentialUseCase) CountBy(column, value, operator string) (res int, err error) {
	repository := repositories.NewCredentialRepository(uc.DB)

	res, err = repository.CountBy(column, value, operator)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

//Function validate credential
func (uc CredentialUseCase) ValidateCredential(column, value, operator, password string) (res bool, err error) {
	//check is data exist
	credential, err := uc.ReadBy(column, value, operator)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	//check password
	res = hashing.CheckHashString(password, credential.Password)

	return res, nil
}
