package request

type SalesRequest struct {
	CustomerID  string               `json:"customer_id"`
	Date        string               `json:"date"`
	Note        string               `json:"note"`
	SalesDetail []SalesDetailRequest `json:"sales_detail"`
}

type SalesDetailRequest struct {
	ProductID string  `json:"product_id"`
	Price     float32 `json:"price"`
	Qty       int     `json:"qty"`
}
