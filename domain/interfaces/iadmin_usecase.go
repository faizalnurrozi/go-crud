package interfaces

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
)

type IAdminUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []view_models.AdminVm, pagination view_models.PaginationVm, err error)

	ReadBy(column, value, operator string) (res view_models.AdminVm, err error)

	Add(req *request.AdminAddRequest) (res string, err error)

	Edit(req *request.AdminEditRequest, ID string) (res string, err error)

	DeleteBy(column, value, operator string) (err error)

	Count(search string) (res int, err error)
}
