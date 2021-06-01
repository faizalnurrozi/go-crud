package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"strings"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(DB *sql.DB) interfaces.IProductRepository {
	return &ProductRepository{DB: DB}
}

func (repository ProductRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Product, err error) {
	model := models.Product{}
	statement := model.GetProductSelectStatement() + ` ` + model.GetProductJoinStatement() + ` ` + model.GetProductWhereStatement() + ` order by p.` + orderBy + ` ` + sort + ` limit $2 offset $3`

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

func (repository ProductRepository) ReadBy(column, value, operator string) (res models.Product, err error) {
	model := models.Product{}
	statement := model.GetProductSelectStatement() + ` ` + model.GetProductJoinStatement() + ` where ` + column + `` + operator + `$1 and p.deleted_at is null`
	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ProductRepository) Add(model models.Product, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO product (merchant_id, sku, name, image, price_po, price_sell, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = tx.QueryRow(statement, model.Merchant.ID, model.Sku, model.Name, model.Image, model.PricePo, model.PriceSell, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ProductRepository) Edit(model models.Product, tx *sql.Tx) (res int64, err error) {
	statement := `update product set merchant_id=$1, sku=$2, name=$3, image=$4, price_po=$5, price_sell=$6, updated_at=$7 where id=$8 and ` + model.GetProductDeleteStatement()
	result, err := tx.Exec(statement, model.Merchant.ID, model.Sku, model.Name, model.Image, model.PricePo, model.PriceSell, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ProductRepository) DeleteBy(column, value, operator string, model models.Product, tx *sql.Tx) (res int64, err error) {
	statement := `update product set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 and ` + model.GetProductDeleteStatement()
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

func (repository ProductRepository) Count(search string) (res int, err error) {
	model := models.Product{}
	statement := `select count(p.id) from product p ` + model.GetProductJoinStatement() + ` ` + model.GetProductWhereStatement()
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
