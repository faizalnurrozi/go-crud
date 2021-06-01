package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type ProductPriceRepository struct {
	DB *sql.DB
}

func NewProductPriceRepository(DB *sql.DB) interfaces.IProductPriceRepository {
	return &ProductPriceRepository{DB: DB}
}

func (repository ProductPriceRepository) ReadBy(outletId, productId string) (res models.ProductPrice, err error) {
	model := models.ProductPrice{}
	statement := model.GetProductPriceSelectStatement() + ` where outlet_id = $1 and product_id = $2`
	row := repository.DB.QueryRow(statement, outletId, productId)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ProductPriceRepository) Add(model models.ProductPrice, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO product_price (outlet_id, product_id, price_po, price_sell, stock) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRow(statement, model.OutletID, model.ProductID, model.PricePo, model.PriceSell, model.Stock).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ProductPriceRepository) Edit(model models.ProductPrice, tx *sql.Tx) (res int64, err error) {
	statement := `update product_price set outlet_id=$1, product_id=$2, price_po=$3, price_sell=$4, stock=$5 where id=$6`
	result, err := tx.Exec(statement, model.OutletID, model.ProductID, model.PricePo, model.PriceSell, model.Stock, model.ID)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ProductPriceRepository) DeleteBy(outletId, productId string, tx *sql.Tx) (res int64, err error) {
	statement := `delete from product_price where outlet_id = $1 and product_id = $2`
	result, err := tx.Exec(statement, outletId, productId)
	if err != nil {
		return res, err
	}
	res, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	return res, nil
}
