package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type SalesDetailRepository struct {
	DB *sql.DB
}

func NewSalesDetailRepository(DB *sql.DB) interfaces.ISalesDetailRepository {
	return &SalesDetailRepository{DB: DB}
}

func (repository SalesDetailRepository) Add(model models.SalesDetail, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO sales_detail (sale_id, product_id, price, qty) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(statement, model.SaleID, model.ProductID, model.Price, model.Qty).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SalesDetailRepository) DeleteBy(saleID string, tx *sql.Tx) (res int64, err error) {
	statement := `delete from sales_detail where sale_id = $1`
	result, err := tx.Exec(statement, saleID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}
