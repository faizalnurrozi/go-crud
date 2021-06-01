package usecase

import (
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
	"github.com/faizalnurrozi/go-crud/repositories"
)

type ProductPriceUseCase struct {
	*Contract
}

func NewProductPriceUseCase(ucContract *Contract) interfaces.IProductPriceUseCase {
	return &ProductPriceUseCase{Contract: ucContract}
}

func (uc ProductPriceUseCase) ReadBy(outletId, productId string) (res view_models.ProductPriceVm, err error) {
	repository := repositories.NewProductPriceRepository(uc.DB)

	productPrice, err := repository.ReadBy(outletId, productId)
	if err != nil {
		fmt.Print("query-readBy: ")
		fmt.Println(err)
		return res, err
	}
	vm := view_models.ProductPriceVm{}
	res = vm.Build(productPrice)

	return res, nil
}

func (uc ProductPriceUseCase) Edit(req *request.ProductPriceRequest, ID string) (res string, err error) {
	repository := repositories.NewProductPriceRepository(uc.DB)

	model := models.ProductPrice{
		ID:        ID,
		OutletID:  req.OutletID,
		ProductID: req.ProductID,
		PricePo:   req.PricePo,
		PriceSell: req.PriceSell,
		Stock:     req.Stock,
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

func (uc ProductPriceUseCase) Add(req *request.ProductPriceRequest) (res string, err error) {
	repository := repositories.NewProductPriceRepository(uc.DB)

	model := models.ProductPrice{
		OutletID:  req.OutletID,
		ProductID: req.ProductID,
		PricePo:   req.PricePo,
		PriceSell: req.PriceSell,
		Stock:     req.Stock,
	}

	res, err = repository.Add(model, uc.TX)
	if err != nil {
		fmt.Print("query-add: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

func (uc ProductPriceUseCase) DeleteBy(outletId, productId string) (err error) {
	repository := repositories.NewProductPriceRepository(uc.DB)

	rowsAffected, err := repository.DeleteBy(outletId, productId, uc.TX)
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
