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

type SalesUseCase struct {
	*Contract
}

func NewSalesUseCase(ucContract *Contract) interfaces.ISalesUseCase {
	return &SalesUseCase{Contract: ucContract}
}

func (uc SalesUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.SalesVm, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewSalesRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.SalesVm{}

	sales, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		fmt.Print("query-browse: ")
		fmt.Println(err)
		return res, pagination, err
	}
	for _, sale := range sales {
		res = append(res, vm.Build(sale))
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

func (uc SalesUseCase) ReadBy(column, value, operator string) (res view_models.SalesVm, err error) {
	repository := repositories.NewSalesRepository(uc.DB)

	sale, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Print("query-readBy: ")
		fmt.Println(err)
		return res, err
	}
	vm := view_models.SalesVm{}
	res = vm.Build(sale)

	return res, nil
}

func (uc SalesUseCase) Edit(req *request.SalesRequest, ID string) (res string, err error) {
	repository := repositories.NewSalesRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Sales{
		ID: ID,
		Customer: models.Customer{
			ID: req.CustomerID,
		},
		Date:      req.Date,
		Note:      req.Note,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = repository.Edit(model, uc.TX)
	if err != nil {
		fmt.Print("query-edit: ")
		fmt.Println(err)
		return res, err
	}

	ucDetail := NewSalesDetailUseCase(uc.Contract)
	err = ucDetail.DeleteBy(ID)
	if err != nil {
		fmt.Print("uc-sale-detail-delete: ")
		fmt.Println(err)
		return res, err
	}

	_, err = ucDetail.Add(req, ID)
	if err != nil {
		fmt.Print("uc-sale-detail-add: ")
		fmt.Println(err)
		return res, err
	}

	res = ID

	return res, nil
}

func (uc SalesUseCase) Add(req *request.SalesRequest) (res string, err error) {
	repository := repositories.NewSalesRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Sales{
		Customer: models.Customer{
			ID: req.CustomerID,
		},
		Date:      req.Date,
		Note:      req.Note,
		CreatedAt: now,
		UpdatedAt: now,
	}

	res, err = repository.Add(model, uc.TX)
	if err != nil {
		fmt.Print("query-add: ")
		fmt.Println(err)
		return res, err
	}

	ucDetail := NewSalesDetailUseCase(uc.Contract)
	_, err = ucDetail.Add(req, res)
	if err != nil {
		fmt.Print("uc-sale-detail-add: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

func (uc SalesUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewSalesRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Sales{
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

func (uc SalesUseCase) Count(search string) (res int, err error) {
	repository := repositories.NewOutletRepository(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		fmt.Print("query-count: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
