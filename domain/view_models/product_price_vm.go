package view_models

import "github.com/faizalnurrozi/go-crud/domain/models"

type ProductPriceVm struct {
	ID        string  `json:"id"`
	PricePo   float32 `json:"price_po"`
	PriceSell float32 `json:"price_sell"`
	Stock     int     `json:"stock"`
}

func (vm ProductPriceVm) Build(model models.ProductPrice) ProductPriceVm {
	return ProductPriceVm{
		ID:        model.ID,
		PricePo:   model.PricePo,
		PriceSell: model.PriceSell,
		Stock:     model.Stock,
	}
}
