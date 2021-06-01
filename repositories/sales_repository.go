package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strings"
)

type SalesRepository struct {
	DB *sql.DB
}

func NewSalesRepository(DB *sql.DB) interfaces.ISalesRepository {
	return &SalesRepository{DB: DB}
}

func (repository SalesRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Sales, err error) {
	model := models.Sales{}
	statement := model.GetSaleSelectStatement() + ` ` + model.GetSaleJoinWithCustomerStatement() + ` ` + model.GetSaleJoinWithDetailStatement() + ` ` + model.GetSaleWhereStatement() + ` ` + model.GetSaleGroupByStatement() + ` order by s.` + orderBy + ` ` + sort + ` limit $2 offset $3`
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

func (repository SalesRepository) ReadBy(column, value, operator string) (res models.Sales, err error) {
	model := models.Sales{}
	statement := model.GetSaleSelectStatement() + ` ` + model.GetSaleJoinWithCustomerStatement() + ` ` + model.GetSaleJoinWithDetailStatement() + ` where ` + column + `` + operator + `$1 and s.deleted_at is null` + ` ` + model.GetSaleGroupByStatement()
	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SalesRepository) Add(model models.Sales, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO sales (customer_id, date, note, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRow(statement, model.Customer.ID, model.Date, model.Note, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SalesRepository) Edit(model models.Sales, tx *sql.Tx) (res int64, err error) {
	statement := `update sales set customer_id=$1, date=$2, note=$3, updated_at=$4 where id=$5 and ` + model.GetSaleDeleteStatement()
	result, err := tx.Exec(statement, model.Customer.ID, model.Date, model.Date, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SalesRepository) DeleteBy(column, value, operator string, model models.Sales, tx *sql.Tx) (res int64, err error) {
	statement := `update sales set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 and ` + model.GetSaleDeleteStatement()
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

func (repository SalesRepository) Count(search string) (res int, err error) {
	model := models.Sales{}
	statement := `select count(s.id) from sales s ` + model.GetSaleSelectStatement() + ` ` + model.GetSaleJoinWithCustomerStatement() + ` ` + model.GetSaleJoinWithDetailStatement() + ` ` + model.GetSaleJoinWithDetailStatement()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
