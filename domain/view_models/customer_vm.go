package view_models

import "github.com/faizalnurrozi/go-crud/domain/models"

type CustomerVm struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

func (vm CustomerVm) Build(model models.Customer) CustomerVm {
	return CustomerVm{
		ID:      model.ID,
		Name:    model.Name,
		Address: model.Address,
		Phone:   model.Phone,
		Email:   model.Email,
	}
}
