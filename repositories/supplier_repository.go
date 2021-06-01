package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strings"
)

type SupplierRepository struct {
	DB *sql.DB
}

func NewSupplierRepository(DB *sql.DB) interfaces.ISupplierRepository {
	return &SupplierRepository{DB: DB}
}

func (repository SupplierRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Supplier, err error) {
	model := models.Supplier{}
	statement := model.GetSupplierSelectStatement() + ` ` + model.GetSupplierWhereStatement() + ` order by s.` + orderBy + ` ` + sort + ` limit $2 offset $3`

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

func (repository SupplierRepository) ReadBy(column, value, operator string) (res models.Supplier, err error) {
	model := models.Supplier{}
	statement := model.GetSupplierSelectStatement() + ` where ` + column + `` + operator + `$1 and s.deleted_at is null`

	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SupplierRepository) Add(model models.Supplier, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO supplier (name, address, phone, email, pic, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = tx.QueryRow(statement, model.Name, model.Address, model.Phone, model.Email, model.Pic, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SupplierRepository) Edit(model models.Supplier, tx *sql.Tx) (res int64, err error) {
	statement := `update supplier set name=$1, address=$2, phone=$3, email=$4, pic=$5, updated_at=$6 where id=$7 and ` + model.GetSupplierDeleteStatement()
	result, err := tx.Exec(statement, model.Name, model.Address, model.Phone, model.Email, model.Pic, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SupplierRepository) DeleteBy(column, value, operator string, model models.Supplier, tx *sql.Tx) (res int64, err error) {
	statement := `update supplier set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 and ` + model.GetSupplierDeleteStatement()
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

func (repository SupplierRepository) Count(search string) (res int, err error) {
	model := models.Supplier{}
	statement := `select count(s.id) from supplier s ` + model.GetSupplierWhereStatement()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
