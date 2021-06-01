package models

import (
	"database/sql"
	"time"
)

type Outlet struct {
	ID        string
	Merchant  Merchant
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const (
	outletSelectStatement    = `select o.id, o.name, m.id, m.name, m.address, m.pic, m.phone, o.created_at, o.updated_at, o.deleted_at from outlet o`
	outletJoinStatement      = `inner join merchant m on o.merchant_id = m.id and m.deleted_at is null`
	outletWhereStatement     = `where lower(o.name) like $1 and m.deleted_at is null`
	outletDeletedAtStatement = `deleted_at is null`
)

func (model Outlet) GetOutletSelectStatement() string {
	return outletSelectStatement
}

func (model Outlet) GetOutletJoinStatement() string {
	return outletJoinStatement
}

func (model Outlet) GetOutletWhereStatement() string {
	return outletWhereStatement
}

func (model Outlet) GetOutletDeleteStatement() string {
	return outletDeletedAtStatement
}

//scan rows for multiple rows
func (model Outlet) ScanRows(rows *sql.Rows) (res Outlet, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.Merchant.ID, &res.Merchant.Name, &res.Merchant.Address, &res.Merchant.Pic, &res.Merchant.Phone, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row for single row
func (model Outlet) ScanRow(row *sql.Row) (res Outlet, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.Merchant.ID, &res.Merchant.Name, &res.Merchant.Address, &res.Merchant.Pic, &res.Merchant.Phone, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
