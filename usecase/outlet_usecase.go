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

type OutletUseCase struct {
	*Contract
}

func NewOutletUseCase(ucContract *Contract) interfaces.IOutletUseCase {
	return &OutletUseCase{Contract: ucContract}
}

func (uc OutletUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.OutletVm, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewOutletRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.OutletVm{}

	outlets, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		fmt.Print("query-browse: ")
		fmt.Println(err)
		return res, pagination, err
	}
	for _, outlet := range outlets {
		res = append(res, vm.Build(outlet))
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

func (uc OutletUseCase) ReadBy(column, value, operator string) (res view_models.OutletVm, err error) {
	repository := repositories.NewOutletRepository(uc.DB)

	outlet, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Print("query-readBy: ")
		fmt.Println(err)
		return res, err
	}
	vm := view_models.OutletVm{}
	res = vm.Build(outlet)

	return res, nil
}

func (uc OutletUseCase) Edit(req *request.OutletRequest, ID string) (res string, err error) {
	repository := repositories.NewOutletRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Outlet{
		ID:   ID,
		Name: req.Name,
		Merchant: models.Merchant{
			ID: req.MerchantID,
		},
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

func (uc OutletUseCase) Add(req *request.OutletRequest) (res string, err error) {
	repository := repositories.NewOutletRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Outlet{
		Name: req.Name,
		Merchant: models.Merchant{
			ID: req.MerchantID,
		},
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

func (uc OutletUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewOutletRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Outlet{
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

func (uc OutletUseCase) Count(search string) (res int, err error) {
	repository := repositories.NewOutletRepository(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		fmt.Print("query-count: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
