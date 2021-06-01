package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
	"github.com/faizalnurrozi/go-crud/repositories"
	"time"
)

type AdminUseCase struct {
	*Contract
}

func NewAdminUseCase(ucContract *Contract) interfaces.IAdminUseCase {
	return &AdminUseCase{Contract: ucContract}
}

func (uc AdminUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.AdminVm, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewAdminRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.AdminVm{}

	//fmt.Println(uc.RoleID)

	//admin browse query
	admins, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		fmt.Println(err)
		return res, pagination, err
	}
	for _, admin := range admins {
		res = append(res, vm.Build(admin))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		fmt.Println(err)
		return res, pagination, err
	}
	pagination = uc.setPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc AdminUseCase) ReadBy(column, value, operator string) (res view_models.AdminVm, err error) {
	repository := repositories.NewAdminRepository(uc.DB)

	admin, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	vm := view_models.AdminVm{}
	res = vm.Build(admin)

	return res, nil
}

func (uc AdminUseCase) Edit(req *request.AdminEditRequest, ID string) (res string, err error) {
	repository := repositories.NewAdminRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Admin{
		ID:        ID,
		Name:      req.Name,
		UpdatedAt: now,
	}
	rowsAffected, err := repository.Edit(model, uc.TX)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	if rowsAffected == 0 {
		return res, errors.New("Row 0")
	}
	res = ID

	return res, nil
}

func (uc AdminUseCase) Add(req *request.AdminAddRequest) (res string, err error) {
	repository := repositories.NewAdminRepository(uc.DB)
	now := time.Now().UTC()

	//add data to credential table
	credentialUc := NewCredentialUseCase(uc.Contract)
	password := req.Password
	credentialID, err := credentialUc.Add(req.Email, password)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	model := models.Admin{
		Name:         req.Name,
		CredentialID: credentialID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	res, err = repository.Add(model, uc.TX)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

func (uc AdminUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewAdminRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Admin{
		UpdatedAt: now,
		DeletedAt: sql.NullTime{Time: now},
	}

	rowsAffected, err := repository.DeleteBy(column, value, operator, model, uc.TX)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Row 0")
	}

	return nil
}

func (uc AdminUseCase) Count(search string) (res int, err error) {
	repository := repositories.NewAdminRepository(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
