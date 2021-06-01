package interfaces

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
)

type ISalesDetailUseCase interface {
	Add(req *request.SalesRequest, saleID string) (res string, err error)
	DeleteBy(saleID string) (err error)
}
