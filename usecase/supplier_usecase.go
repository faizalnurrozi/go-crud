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

type SupplierUseCase struct {
	*Contract
}

func NewSupplierUseCase(ucContract *Contract) interfaces.ISupplierUseCase {
	return &SupplierUseCase{Contract: ucContract}
}

func (uc SupplierUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.SupplierVm, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewSupplierRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.SupplierVm{}

	suppliers, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		fmt.Print("query-browse: ")
		fmt.Println(err)
		return res, pagination, err
	}
	for _, supplier := range suppliers {
		res = append(res, vm.Build(supplier))
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

func (uc SupplierUseCase) ReadBy(column, value, operator string) (res view_models.SupplierVm, err error) {
	repository := repositories.NewSupplierRepository(uc.DB)

	supplier, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Print("query-readBy: ")
		fmt.Println(err)
		return res, err
	}
	vm := view_models.SupplierVm{}
	res = vm.Build(supplier)

	return res, nil
}

func (uc SupplierUseCase) Edit(req *request.SupplierRequest, ID string) (res string, err error) {
	repository := repositories.NewSupplierRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Supplier{
		ID:        ID,
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     req.Email,
		Pic:       req.Pic,
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

func (uc SupplierUseCase) Add(req *request.SupplierRequest) (res string, err error) {
	repository := repositories.NewSupplierRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Supplier{
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     req.Email,
		Pic:       req.Pic,
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

func (uc SupplierUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewSupplierRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Supplier{
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

func (uc SupplierUseCase) Count(search string) (res int, err error) {
	repository := repositories.NewSupplierRepository(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		fmt.Print("query-count: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
