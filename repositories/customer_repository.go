package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strings"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(DB *sql.DB) interfaces.ICustomerRepository {
	return &CustomerRepository{DB: DB}
}

func (repository CustomerRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Customer, err error) {
	model := models.Customer{}
	statement := model.GetCustomerSelectStatement() + ` ` + model.GetCustomerWhereStatement() + ` order by c.` + orderBy + ` ` + sort + ` limit $2 offset $3`

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

func (repository CustomerRepository) ReadBy(column, value, operator string) (res models.Customer, err error) {
	model := models.Customer{}
	statement := model.GetCustomerSelectStatement() + ` where ` + column + `` + operator + `$1 and c.deleted_at is null`

	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository CustomerRepository) Add(model models.Customer, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO customer (name, address, phone, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = tx.QueryRow(statement, model.Name, model.Address, model.Phone, model.Email, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository CustomerRepository) Edit(model models.Customer, tx *sql.Tx) (res int64, err error) {
	statement := `update customer set name=$1, address=$2, phone=$3, email=$4, updated_at=$5 where id=$6 and ` + model.GetCustomerDeleteStatement()
	result, err := tx.Exec(statement, model.Name, model.Address, model.Phone, model.Email, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository CustomerRepository) DeleteBy(column, value, operator string, model models.Customer, tx *sql.Tx) (res int64, err error) {
	statement := `update customer set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 and ` + model.GetCustomerDeleteStatement()
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

func (repository CustomerRepository) Count(search string) (res int, err error) {
	model := models.Customer{}
	statement := `select count(c.id) from customer c ` + model.GetCustomerWhereStatement()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
