package usecase

import (
	"database/sql"
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
	"github.com/faizalnurrozi/go-crud/repositories"
	"time"
)

type CustomerUseCase struct {
	*Contract
}

func NewCustomerUseCase(ucContract *Contract) interfaces.ICustomerUseCase {
	return &CustomerUseCase{Contract: ucContract}
}

func (uc CustomerUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.CustomerVm, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewCustomerRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.CustomerVm{}

	customers, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		fmt.Print("query-browse: ")
		fmt.Println(err)
		return res, pagination, err
	}
	for _, customer := range customers {
		res = append(res, vm.Build(customer))
	}

	totalCount, err := uc.Count(search)
	if err != nil {
		fmt.Print("uc-count: ")
		fmt.Println(err)
		return res, pagination, err
	}
	pagination = uc.setPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CustomerUseCase) ReadBy(column, value, operator string) (res view_models.CustomerVm, err error) {
	repository := repositories.NewCustomerRepository(uc.DB)

	customer, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Print("query-readBy: ")
		fmt.Println(err)
		return res, err
	}
	vm := view_models.CustomerVm{}
	res = vm.Build(customer)

	return res, nil
}

func (uc CustomerUseCase) Edit(req *request.CustomerRequest, ID string) (res string, err error) {
	repository := repositories.NewCustomerRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Customer{
		ID:        ID,
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     req.Email,
		UpdatedAt: now,
	}
	rowsAffected, err := repository.Edit(model, uc.TX)
	if err != nil {
		fmt.Print("query-Edit: ")
		fmt.Println(err)
		return res, err
	}
	if rowsAffected == 0 {
		return res, err
	}
	res = ID

	return res, nil
}

func (uc CustomerUseCase) Add(req *request.CustomerRequest) (res string, err error) {
	repository := repositories.NewCustomerRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Customer{
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
	res, err = repository.Add(model, uc.TX)
	if err != nil {
		fmt.Print("query-add: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

func (uc CustomerUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewCustomerRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Customer{
		UpdatedAt: now,
		DeletedAt: sql.NullTime{Time: now},
	}

	rowsAffected, err := repository.DeleteBy(column, value, operator, model, uc.TX)
	if err != nil {
		fmt.Print("query-deleteBy: ")
		fmt.Println(err)
		return err
	}
	if rowsAffected == 0 {
		return err
	}

	return nil
}

func (uc CustomerUseCase) Count(search string) (res int, err error) {
	repository := repositories.NewCustomerRepository(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		fmt.Print("query-count: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
