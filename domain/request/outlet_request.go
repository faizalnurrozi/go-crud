package request

type OutletRequest struct {
	Name       string `json:"name" validate:"required"`
	MerchantID string `json:"merchant_id"`
}
