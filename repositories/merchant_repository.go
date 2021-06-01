package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strings"
)

type MerhcantRepository struct {
	DB *sql.DB
}

func NewMerhcantRepository(DB *sql.DB) interfaces.IMerchantRepository {
	return &MerhcantRepository{DB: DB}
}

func (repository MerhcantRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Merchant, err error) {
	model := models.Merchant{}
	statement := model.GetMerchantSelectStatement() + ` ` + model.GetMerchantWhereStatement() + ` order by m.` + orderBy + ` ` + sort + ` limit $2 offset $3`

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

func (repository MerhcantRepository) ReadBy(column, value, operator string) (res models.Merchant, err error) {
	model := models.Merchant{}
	statement := model.GetMerchantSelectStatement() + ` where ` + column + `` + operator + `$1 and m.deleted_at is null`

	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MerhcantRepository) Add(model models.Merchant, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO merchant (name, address, pic, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = tx.QueryRow(statement, model.Name, model.Address, model.Pic, model.Phone, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MerhcantRepository) Edit(model models.Merchant, tx *sql.Tx) (res int64, err error) {
	statement := `update merchant set name=$1, address=$2, pic=$3, phone=$4, updated_at=$5 where id=$6 and ` + model.GetMerchantDeleteStatement()
	result, err := tx.Exec(statement, model.Name, model.Address, model.Pic, model.Phone, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MerhcantRepository) DeleteBy(column, value, operator string, model models.Merchant, tx *sql.Tx) (res int64, err error) {
	statement := `update merchant set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 and ` + model.GetMerchantDeleteStatement()
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

func (repository MerhcantRepository) Count(search string) (res int, err error) {
	model := models.Merchant{}
	statement := `select count(m.id) from merchant m ` + model.GetMerchantWhereStatement()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
