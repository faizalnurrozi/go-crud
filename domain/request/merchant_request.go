package request

type MerchantRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Pic     string `json:"pic" validate:"required"`
}
