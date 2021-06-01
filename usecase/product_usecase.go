package usecase

import (
	"database/sql"
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
	"github.com/faizalnurrozi/go-crud/repositories"
	"strconv"
	"strings"
	"time"
)

type ProductUseCase struct {
	*Contract
}

func NewProductUseCase(ucContract *Contract) interfaces.IProductUseCase {
	return &ProductUseCase{Contract: ucContract}
}

func (uc ProductUseCase) Browse(search, orderBy, sort string, page, limit int) (res []view_models.ProductVm, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewProductRepository(uc.DB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)
	vm := view_models.ProductVm{}

	products, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		fmt.Print("query-browse: ")
		fmt.Println(err)
		return res, pagination, err
	}
	for _, product := range products {
		res = append(res, vm.Build(product))
	}

	totalCount, err := uc.Count(search)
	if err != nil {
		fmt.Print("uc-count: ")
		fmt.Println(err)
		return res, pagination, err
	}
	pagination = uc.setPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc ProductUseCase) ReadBy(column, value, operator string) (res view_models.ProductVm, err error) {
	repository := repositories.NewProductRepository(uc.DB)

	product, err := repository.ReadBy(column, value, operator)
	if err != nil {
		fmt.Print("query-readBy: ")
		fmt.Println(err)
		return res, err
	}
	vm := view_models.ProductVm{}
	res = vm.Build(product)

	return res, nil
}

func (uc ProductUseCase) Edit(input map[string]interface{}, ID string) (res string, err error) {
	repository := repositories.NewProductRepository(uc.DB)
	now := time.Now().UTC()

	sku, _ := input["sku"]
	name, _ := input["name"]
	image, _ := input["image"]
	pricePoParam, _ := input["pricePo"]
	priceSellParam, _ := input["priceSell"]
	merchantID, _ := input["merchantID"]

	pricePo, err := strconv.ParseFloat(pricePoParam.(string), 32)
	if err != nil {
		fmt.Println(err)
	}
	priceSell, err := strconv.ParseFloat(priceSellParam.(string), 32)
	if err != nil {
		fmt.Println(err)
	}

	model := models.Product{
		ID:        ID,
		Sku:       strings.ToLower(sku.(string)),
		Name:      strings.ToLower(name.(string)),
		Image:     strings.ToLower(image.(string)),
		PricePo:   float32(pricePo),
		PriceSell: float32(priceSell),
		Merchant: models.Merchant{
			ID: strings.ToLower(merchantID.(string)),
		},
		UpdatedAt: now,
	}
	rowsAffected, err := repository.Edit(model, uc.TX)
	if err != nil {
		fmt.Print("query-Edit: ")
		fmt.Println(err)
		return res, err
	}
	if rowsAffected == 0 {
		return res, err
	}
	res = ID

	return res, nil
}

func (uc ProductUseCase) Add(input map[string]interface{}) (res string, err error) {
	repository := repositories.NewProductRepository(uc.DB)
	now := time.Now().UTC()

	sku, _ := input["sku"]
	name, _ := input["name"]
	image, _ := input["image"]
	pricePoParam, _ := input["pricePo"]
	priceSellParam, _ := input["priceSell"]
	merchantID, _ := input["merchantID"]

	pricePo, err := strconv.ParseFloat(pricePoParam.(string), 32)
	if err != nil {
		fmt.Println(err)
	}
	priceSell, err := strconv.ParseFloat(priceSellParam.(string), 32)
	if err != nil {
		fmt.Println(err)
	}

	model := models.Product{
		Sku:       strings.ToLower(sku.(string)),
		Name:      strings.ToLower(name.(string)),
		Image:     strings.ToLower(image.(string)),
		PricePo:   float32(pricePo),
		PriceSell: float32(priceSell),
		Merchant: models.Merchant{
			ID: strings.ToLower(merchantID.(string)),
		},
		UpdatedAt: now,
	}

	res, err = repository.Add(model, uc.TX)
	if err != nil {
		fmt.Print("query-add: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}

func (uc ProductUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := repositories.NewProductRepository(uc.DB)
	now := time.Now().UTC()

	model := models.Product{
		UpdatedAt: now,
		DeletedAt: sql.NullTime{Time: now},
	}

	rowsAffected, err := repository.DeleteBy(column, value, operator, model, uc.TX)
	if err != nil {
		fmt.Print("query-deleteBy: ")
		fmt.Println(err)
		return err
	}
	if rowsAffected == 0 {
		return err
	}

	return nil
}

func (uc ProductUseCase) Count(search string) (res int, err error) {
	repository := repositories.NewProductRepository(uc.DB)
	res, err = repository.Count(search)
	if err != nil {
		fmt.Print("query-count: ")
		fmt.Println(err)
		return res, err
	}

	return res, nil
}
