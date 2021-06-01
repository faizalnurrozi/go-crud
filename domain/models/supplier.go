package models

import (
	"database/sql"
	"time"
)

type Supplier struct {
	ID        string
	Name      string
	Address   string
	Phone     string
	Email     string
	Pic       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const (
	supplierSelectStatement    = `select s.id, s.name, s.address, s.phone, s.email, s.pic, s.created_at, s.updated_at, s.deleted_at from supplier s`
	supplierWhereStatement     = `where (lower(s.name) like $1 or lower(s.email) like $1) and s.deleted_at is null`
	supplierDeletedAtStatement = `deleted_at is null`
)

func (model Supplier) GetSupplierSelectStatement() string {
	return supplierSelectStatement
}

func (model Supplier) GetSupplierWhereStatement() string {
	return supplierWhereStatement
}

func (model Supplier) GetSupplierDeleteStatement() string {
	return supplierDeletedAtStatement
}

//scan rows for multiple rows
func (model Supplier) ScanRows(rows *sql.Rows) (res Supplier, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.Address, &res.Phone, &res.Email, &res.Pic, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row for single row
func (model Supplier) ScanRow(row *sql.Row) (res Supplier, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.Address, &res.Phone, &res.Email, &res.Pic, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
