package request

type ProductPriceRequest struct {
	OutletID  string  `json:"outlet_id" validate:"required"`
	ProductID string  `json:"product_id" validate:"required"`
	PricePo   float32 `json:"price_po" validate:"required"`
	PriceSell float32 `json:"price_sell" validate:"required"`
	Stock     int     `json:"stock" validate:"required"`
}
