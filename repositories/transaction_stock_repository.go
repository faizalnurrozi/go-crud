package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type TransactionStockRepository struct {
	DB *sql.DB
}

func NewTransactionStockRepository(DB *sql.DB) interfaces.ITransactionStockRepository {
	return &TransactionStockRepository{DB: DB}
}

func (repository TransactionStockRepository) Add(model models.TransactionStock, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO transaction_stock (transaction_id, date, product_id, merchant_id, stock_in, stock_in) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = tx.QueryRow(statement, model.TransactionID, model.Date, model.ProductID, model.MerchantID, model.StockIn, model.StockOut).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository TransactionStockRepository) Diff(merchantId, outletId, productId string) (res int, err error) {
	statement := `select SUM(stock_in)-SUM(stock_out) from transaction_stock ts where merchant_id = $1 and outlet_it = $2 and product_id = $3`
	err = repository.DB.QueryRow(statement, merchantId, outletId, productId).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
