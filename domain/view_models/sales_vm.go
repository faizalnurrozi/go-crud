package view_models

import (
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strconv"
	"strings"
)

type SalesVm struct {
	ID          string          `json:"id"`
	Customer    CustomerVm      `json:"customer"`
	Date        string          `json:"date"`
	Note        string          `json:"note"`
	SalesDetail []SalesDetailVm `json:"sales_detail"`
}

type SalesDetailVm struct {
	Product ProductVm `json:"product"`
	Price   float32   `json:"price"`
	Qty     int       `json:"qty"`
}

func (vm SalesVm) Build(model models.Sales) SalesVm {

	var salesDetailVm []SalesDetailVm

	productID := strings.Split(model.ProductID, ",")
	productSku := strings.Split(model.ProductSku, ",")
	productName := strings.Split(model.ProductName, ",")
	sdPrice := strings.Split(model.SDPrice, ",")
	sdQty := strings.Split(model.SDQty, ",")

	for index, sku := range productSku {

		price, _ := strconv.ParseFloat(sdPrice[index], 32)
		qty, _ := strconv.Atoi(sdQty[index])

		salesDetailVm = append(salesDetailVm, SalesDetailVm{
			Product: ProductVm{
				ID:   productID[index],
				Sku:  sku,
				Name: productName[index],
			},
			Price: float32(price),
			Qty:   qty,
		})
	}

	saleVm := SalesVm{
		ID: model.ID,
		Customer: CustomerVm{
			ID:      model.Customer.ID,
			Name:    model.Customer.Name,
			Address: model.Customer.Address,
			Phone:   model.Customer.Phone,
			Email:   model.Customer.Email,
		},
		Date:        model.Date,
		Note:        model.Date,
		SalesDetail: salesDetailVm,
	}

	return saleVm
}
