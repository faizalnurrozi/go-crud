package models

import (
	"database/sql"
	"time"
)

type Credential struct {
	ID        string
	Email     string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const (
	CredentialSelectStatement = `select c.id,c.email,c.password,c.created_at,c.updated_at from credentials c`
	CredentialWhereStatement  = `where c.deleted_at is null`
)

func NewCredential() Credential {
	return Credential{}
}

//Scan for single row
func (model Credential) ScanRow(row *sql.Row) (res Credential, err error) {
	err = row.Scan(&res.ID, &res.Email, &res.Password, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//Scan for multiple rows
func (model Credential) ScanRows(rows *sql.Row) (res Credential, err error) {
	err = rows.Scan(&res.ID, &res.Email, &res.Password, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
