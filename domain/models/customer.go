package models

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID        string
	Name      string
	Address   string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const (
	customerSelectStatement    = `select c.id, c.name, c.address, c.phone, c.email, c.created_at, c.updated_at, c.deleted_at from customer c`
	customerWhereStatement     = `where (lower(c.name) like $1 or lower(c.email) like $1) and c.deleted_at is null`
	customerDeletedAtStatement = `deleted_at is null`
)

func (model Customer) GetCustomerSelectStatement() string {
	return customerSelectStatement
}

func (model Customer) GetCustomerWhereStatement() string {
	return customerWhereStatement
}

func (model Customer) GetCustomerDeleteStatement() string {
	return customerDeletedAtStatement
}

//scan rows for multiple rows
func (model Customer) ScanRows(rows *sql.Rows) (res Customer, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.Address, &res.Phone, &res.Email, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row for single row
func (model Customer) ScanRow(row *sql.Row) (res Customer, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.Address, &res.Phone, &res.Email, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
