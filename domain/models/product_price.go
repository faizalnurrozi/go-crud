package models

import (
	"database/sql"
)

type ProductPrice struct {
	ID        string
	OutletID  string
	ProductID string
	PricePo   float32
	PriceSell float32
	Stock     int
}

const (
	productPriceSelectStatement = `select pp.id, pp.outlet_id, pp.product_id, pp.price_po, pp.price_sell, pp.stock from product_price pp`
	productPriceWhereStatement  = `where lower(pp.outlet_id) like $1 and lower(pp.product_id) like $2`
)

func (model ProductPrice) GetProductPriceSelectStatement() string {
	return productPriceSelectStatement
}

func (model ProductPrice) GetProductPriceWhereStatement() string {
	return productPriceWhereStatement
}

//scan row for single row
func (model ProductPrice) ScanRow(row *sql.Row) (res ProductPrice, err error) {
	err = row.Scan(&res.ID, &res.OutletID, &res.ProductID, &res.PricePo, &res.PriceSell, &res.Stock)
	if err != nil {
		return res, err
	}

	return res, nil
}
