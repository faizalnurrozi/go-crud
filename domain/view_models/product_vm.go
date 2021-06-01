package view_models

import "github.com/faizalnurrozi/go-crud/domain/models"

type ProductVm struct {
	ID        string     `json:"id"`
	Sku       string     `json:"sku"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	PricePo   float32    `json:"price_po"`
	PriceSell float32    `json:"price_sell"`
	Merchant  MerchantVm `json:"merchant"`
}

func (vm ProductVm) Build(model models.Product) ProductVm {
	return ProductVm{
		ID:        model.ID,
		Sku:       model.Sku,
		Name:      model.Name,
		Image:     model.Image,
		PricePo:   model.PricePo,
		PriceSell: vm.PriceSell,
		Merchant: MerchantVm{
			ID:      model.Merchant.ID,
			Name:    model.Merchant.Name,
			Address: model.Merchant.Address,
			Pic:     model.Merchant.Pic,
			Phone:   model.Merchant.Phone,
		},
	}
}
