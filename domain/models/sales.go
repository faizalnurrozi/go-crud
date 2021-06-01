package models

import (
	"database/sql"
	"time"
)

type Sales struct {
	ID          string
	Customer    Customer
	Date        string
	Note        string
	ProductID   string
	ProductSku  string
	ProductName string
	SDPrice     string
	SDQty       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

const (
	saleSelectStatement           = `select s.id, s.date, s.note, c.id, c.name, c.address, c.phone, c.email, array_to_string(array_agg(p.id), ',') sd_product_id, array_to_string(array_agg(p.sku), ',') sd_product_sku, array_to_string(array_agg(p.name), ',') sd_product_name, array_to_string(array_agg(sd.price), ',') sd_price, array_to_string(array_agg(sd.qty), ',') sd_qty, s.created_at, s.updated_at, s.deleted_at from sales s`
	saleJoinWithCustomerStatement = `inner join customer c on s.customer_id = c.id and s.deleted_at is null`
	saleJoinWithDetailStatement   = `inner join sales_detail sd on s.id = sd.sale_id`
	saleJoinWithProductStatement  = `inner join product p on sd.product_id = p.id`
	saleWhereStatement            = `where (lower(c.name) like $1 or lower(s.note) like $1) and s.deleted_at is null`
	saleGroupByStatement          = `group by s.id, c.id`
	saleDeletedAtStatement        = `deleted_at is null`
)

func (model Sales) GetSaleSelectStatement() string {
	return saleSelectStatement
}

func (model Sales) GetSaleJoinWithCustomerStatement() string {
	return saleJoinWithCustomerStatement
}

func (model Sales) GetSaleJoinWithDetailStatement() string {
	return saleJoinWithDetailStatement + ` ` + saleJoinWithProductStatement
}

func (model Sales) GetSaleWhereStatement() string {
	return saleWhereStatement
}

func (model Sales) GetSaleGroupByStatement() string {
	return saleGroupByStatement
}

func (model Sales) GetSaleDeleteStatement() string {
	return saleDeletedAtStatement
}

//scan rows for multiple rows
func (model Sales) ScanRows(rows *sql.Rows) (res Sales, err error) {
	err = rows.Scan(&res.ID, &res.Date, &res.Note, &res.Customer.ID, &res.Customer.Name, &res.Customer.Address, &res.Customer.Phone, &res.Customer.Email, &res.ProductID, &res.ProductSku, &res.ProductName, &res.SDPrice, &res.SDQty, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row for single row
func (model Sales) ScanRow(row *sql.Row) (res Sales, err error) {
	err = row.Scan(&res.ID, &res.Date, &res.Note, &res.Customer.ID, &res.Customer.Name, &res.Customer.Address, &res.Customer.Phone, &res.Customer.Email, &res.ProductID, &res.ProductSku, &res.ProductName, &res.SDPrice, &res.SDQty, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
