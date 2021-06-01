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

type MerchantUseCase struct {
	*Contract
}

func NewMerchantUseCase(ucContract *Contract) interfaces.IMerchantUseCase {
	return &MerchantUseCase{Contract: ucContract}
}

func (uc MerchantUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.MerchantVm, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewMerhcantRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.MerchantVm{}

	merchants, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		fmt.Print("query-browse: ")
		fmt.Println(err)
		return res, pagination, err
	}
	for _, merchant := range merchants {
		res = append(res, vm.Build(merchant))
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

func (uc MerchantUseCase) ReadBy(column, value, operator string) (res view_models.MerchantVm, err error) {
	repository := repositories.NewMerhcantRepository(uc.DB)

	merchant, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Print("query-readBy: ")
		fmt.Println(err)
		return res, err
	}
	vm := view_models.MerchantVm{}
	res = vm.Build(merchant)

	return res, nil
}

func (uc MerchantUseCase) Edit(req *request.MerchantRequest, ID string) (res string, err error) {
	repository := repositories.NewMerhcantRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Merchant{
		ID:        ID,
		Name:      req.Name,
		Address:   req.Address,
		Pic:       req.Pic,
		Phone:     req.Phone,
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

func (uc MerchantUseCase) Add(req *request.MerchantRequest) (res string, err error) {
	repository := repositories.NewMerhcantRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Merchant{
		Name:      req.Name,
		Address:   req.Address,
		Pic:       req.Pic,
		Phone:     req.Phone,
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

func (uc MerchantUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewMerhcantRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Merchant{
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

func (uc MerchantUseCase) Count(search string) (res int, err error) {
	repository := repositories.NewMerhcantRepository(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		fmt.Print("query-count: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
