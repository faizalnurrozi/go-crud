package interfaces

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
)

type IProductPriceUseCase interface {
	ReadBy(outletId, productId string) (res view_models.ProductPriceVm, err error)

	Add(req *request.ProductPriceRequest) (res string, err error)

	Edit(req *request.ProductPriceRequest, ID string) (res string, err error)

	DeleteBy(outletId, productId string) (err error)
}
