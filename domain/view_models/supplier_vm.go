package view_models

import "github.com/faizalnurrozi/go-crud/domain/models"

type SupplierVm struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Pic     string `json:"pic"`
}

func (vm SupplierVm) Build(model models.Supplier) SupplierVm {
	return SupplierVm{
		ID:      model.ID,
		Name:    model.Name,
		Address: model.Address,
		Phone:   model.Phone,
		Email:   model.Email,
		Pic:     model.Pic,
	}
}
