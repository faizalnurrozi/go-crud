package models

import (
	"database/sql"
	"time"
)

type Merchant struct {
	ID        string
	Name      string
	Address   string
	Pic       string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const (
	merchantSelectStatement    = `select m.id, m.name, m.address, m.pic, m.phone, m.created_at, m.updated_at, m.deleted_at from merchant m`
	merchantWhereStatement     = `where (lower(m.name) like $1 or lower(m.pic) like $1) and m.deleted_at is null`
	merchantDeletedAtStatement = `deleted_at is null`
)

func (model Merchant) GetMerchantSelectStatement() string {
	return merchantSelectStatement
}

func (model Merchant) GetMerchantWhereStatement() string {
	return merchantWhereStatement
}

func (model Merchant) GetMerchantDeleteStatement() string {
	return merchantDeletedAtStatement
}

//scan rows for multiple rows
func (model Merchant) ScanRows(rows *sql.Rows) (res Merchant, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.Address, &res.Pic, &res.Phone, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row for single row
func (model Merchant) ScanRow(row *sql.Row) (res Merchant, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.Address, &res.Pic, &res.Phone, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
