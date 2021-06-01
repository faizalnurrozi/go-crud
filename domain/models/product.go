package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        string
	Merchant  Merchant
	Sku       string
	Name      string
	Image     string
	PricePo   float32
	PriceSell float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const (
	productSelectStatement    = `select p.id, p.sku, p.name, p.image, p.price_po, p.price_sell, m.id, m.name, m.address, m.pic, m.phone, p.created_at, p.updated_at, p.deleted_at from product p`
	productJoinStatement      = `inner join merchant m on p.merchant_id = m.id and m.deleted_at is null`
	productWhereStatement     = `where (lower(p.name) like $1 or lower(p.sku) like $1) and p.deleted_at is null`
	productDeletedAtStatement = `deleted_at is null`
)

func (model Product) GetProductSelectStatement() string {
	return productSelectStatement
}

func (model Product) GetProductJoinStatement() string {
	return productJoinStatement
}

func (model Product) GetProductWhereStatement() string {
	return productWhereStatement
}

func (model Product) GetProductDeleteStatement() string {
	return productDeletedAtStatement
}

//scan rows for multiple rows
func (model Product) ScanRows(rows *sql.Rows) (res Product, err error) {
	err = rows.Scan(&res.ID, &res.Sku, &res.Name, &res.Image, &res.PricePo, &res.PriceSell, &res.Merchant.ID, &res.Merchant.Name, &res.Merchant.Address, &res.Merchant.Pic, &res.Merchant.Phone, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row for single row
func (model Product) ScanRow(row *sql.Row) (res Product, err error) {
	err = row.Scan(&res.ID, &res.Sku, &res.Name, &res.Image, &res.PricePo, &res.PriceSell, &res.Merchant.ID, &res.Merchant.Name, &res.Merchant.Address, &res.Merchant.Pic, &res.Merchant.Phone, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
