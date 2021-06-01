package interfaces

import (
	"github.com/faizalnurrozi/go-crud/domain/view_models"
)

type IProductUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []view_models.ProductVm, pagination view_models.PaginationVm, err error)

	ReadBy(column, value, operator string) (res view_models.ProductVm, err error)

	Add(input map[string]interface{}) (res string, err error)

	Edit(input map[string]interface{}, ID string) (res string, err error)

	DeleteBy(column, value, operator string) (err error)

	Count(search string) (res int, err error)
}
