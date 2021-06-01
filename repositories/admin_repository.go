package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strings"
)

type AdminRepository struct {
	DB *sql.DB
}

func NewAdminRepository(DB *sql.DB) interfaces.IAdminRepository {
	return &AdminRepository{DB: DB}
}

func (repository AdminRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Admin, err error) {
	model := models.Admin{}
	statement := model.GetAdminSelectStatement() + ` ` + model.GetAdminJoinStatement() + ` ` + model.GetAdminWhereStatement() + ` order by a.` + orderBy + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := model.ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil
}

func (repository AdminRepository) ReadBy(column, value, operator string) (res models.Admin, err error) {
	model := models.Admin{}
	statement := model.GetAdminSelectStatement() + ` ` + model.GetAdminJoinStatement() + ` WHERE ` + column + `` + operator + `$1 AND a.deleted_at is null`

	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

//TODO: Data already exist when deleted_at is not null
func (repository AdminRepository) Add(model models.Admin, tx *sql.Tx) (res string, err error) {
	statement := `insert into admins (name,credential_id,created_at,updated_at) values ($1,$2,$3,$4) returning id`
	err = tx.QueryRow(statement, model.Name, model.CredentialID, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

//TODO: Edit role_id + email
func (repository AdminRepository) Edit(model models.Admin, tx *sql.Tx) (res int64, err error) {
	statement := `update admins set name=$1, updated_at=$2 where id=$3 and ` + model.GetAdminDeletedAtStatement()
	result, err := tx.Exec(statement, model.Name, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository AdminRepository) DeleteBy(column, value, operator string, model models.Admin, tx *sql.Tx) (res int64, err error) {
	statement := `update admins set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 and ` + model.GetAdminDeletedAtStatement()
	result, err := tx.Exec(statement, model.UpdatedAt, model.DeletedAt.Time, value)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository AdminRepository) Count(search string) (res int, err error) {
	model := models.Admin{}
	statement := `select count (a.id) from admins a ` + model.GetAdminJoinStatement() + ` ` + model.GetAdminWhereStatement()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
