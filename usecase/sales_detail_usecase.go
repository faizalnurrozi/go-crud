package usecase

import (
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/repositories"
)

type SalesDetailUseCase struct {
	*Contract
}

func NewSalesDetailUseCase(ucContract *Contract) interfaces.ISalesDetailUseCase {
	return &SalesDetailUseCase{Contract: ucContract}
}

func (uc SalesDetailUseCase) Add(req *request.SalesRequest, saleID string) (res string, err error) {
	repository := repositories.NewSalesDetailRepository(uc.DB)

	for _, detail := range req.SalesDetail {
		model := models.SalesDetail{
			SaleID:    saleID,
			ProductID: detail.ProductID,
			Price:     detail.Price,
			Qty:       detail.Qty,
		}

		res, err = repository.Add(model, uc.TX)
		if err != nil {
			fmt.Print("query-add: ")
			fmt.Println(err)
			return res, err
		}

	}

	return res, nil
}

func (uc SalesDetailUseCase) DeleteBy(saleID string) (err error) {
	repository := repositories.NewSalesDetailRepository(uc.DB)

	rowsAffected, err := repository.DeleteBy(saleID, uc.TX)
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
