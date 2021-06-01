package view_models

import "github.com/faizalnurrozi/go-crud/domain/models"

type OutletVm struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Merchant MerchantVm `json:"merchant"`
}

func (vm OutletVm) Build(model models.Outlet) OutletVm {
	return OutletVm{
		ID:   model.ID,
		Name: model.Name,
		Merchant: MerchantVm{
			ID:      model.Merchant.ID,
			Name:    model.Merchant.Name,
			Address: model.Merchant.Address,
			Pic:     model.Merchant.Pic,
			Phone:   model.Merchant.Phone,
		},
	}
}
