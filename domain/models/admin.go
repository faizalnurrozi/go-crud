package models

import (
	"database/sql"
	"time"
)

type Admin struct {
	ID           string
	Name         string
	CredentialID string
	Credential   Credential
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}

const (
	adminSelectStatement    = `select a.id,a.name,a.credential_id,a.created_at,a.updated_at,a.deleted_at,c.email,c.password from admins a`
	adminJoinStatement      = `inner join credentials c on c.id = a.credential_id and c.deleted_at is null`
	adminWhereStatement     = `where (lower(a.name) like $1 or lower(c.email) like $1) and a.deleted_at is null`
	adminDeletedAtStatement = `deleted_at IS NULL`
)

func (model Admin) GetAdminSelectStatement() string {
	return adminSelectStatement
}

func (model Admin) GetAdminJoinStatement() string {
	return adminJoinStatement
}

func (model Admin) GetAdminWhereStatement() string {
	return adminWhereStatement
}

func (model Admin) GetAdminDeletedAtStatement() string {
	return adminDeletedAtStatement
}

//scan rows for multiple rows
func (model Admin) ScanRows(rows *sql.Rows) (res Admin, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.CredentialID, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.Credential.Email, &res.Credential.Password)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row for single row
func (model Admin) ScanRow(row *sql.Row) (res Admin, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.CredentialID, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.Credential.Email, &res.Credential.Password)
	if err != nil {
		return res, err
	}

	return res, nil
}
