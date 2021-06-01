package view_models

import "github.com/faizalnurrozi/go-crud/domain/models"

type MerchantVm struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Pic     string `json:"pic"`
	Phone   string `json:"phone"`
}

func (vm MerchantVm) Build(model models.Merchant) MerchantVm {
	return MerchantVm{
		ID:      model.ID,
		Name:    model.Name,
		Address: model.Address,
		Phone:   model.Phone,
		Pic:     model.Pic,
	}
}
