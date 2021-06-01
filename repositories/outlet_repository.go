package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strings"
)

type OutletRepository struct {
	DB *sql.DB
}

func NewOutletRepository(DB *sql.DB) interfaces.IOutletRepository {
	return &OutletRepository{DB: DB}
}

func (repository OutletRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Outlet, err error) {
	model := models.Outlet{}
	statement := model.GetOutletSelectStatement() + ` ` + model.GetOutletJoinStatement() + ` ` + model.GetOutletWhereStatement() + ` order by o.` + orderBy + ` ` + sort + ` limit $2 offset $3`

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

func (repository OutletRepository) ReadBy(column, value, operator string) (res models.Outlet, err error) {
	model := models.Outlet{}
	statement := model.GetOutletSelectStatement() + ` ` + model.GetOutletJoinStatement() + ` where ` + column + `` + operator + `$1 and o.deleted_at is null`
	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository OutletRepository) Add(model models.Outlet, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO outlet (name, merchant_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(statement, model.Name, model.Merchant.ID, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository OutletRepository) Edit(model models.Outlet, tx *sql.Tx) (res int64, err error) {
	statement := `update outlet set name=$1, merchant_id=$2, updated_at=$3 where id=$4 and ` + model.GetOutletDeleteStatement()
	result, err := tx.Exec(statement, model.Name, model.Merchant.ID, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository OutletRepository) DeleteBy(column, value, operator string, model models.Outlet, tx *sql.Tx) (res int64, err error) {
	statement := `update outlet set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 and ` + model.GetOutletDeleteStatement()
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

func (repository OutletRepository) Count(search string) (res int, err error) {
	model := models.Outlet{}
	statement := `select count(o.id) from outlet o ` + model.GetOutletJoinStatement() + ` ` + model.GetOutletWhereStatement()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
