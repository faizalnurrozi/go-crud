package handlers

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductPriceHandler struct {
	Handler
}

func (handler ProductPriceHandler) ReadByID(ctx *fiber.Ctx) (err error) {
	// Get Param
	outletId := ctx.Query("outlet_id")
	productId := ctx.Query("product_id")

	// Database Processing
	uc := usecase.NewProductPriceUseCase(handler.UcContract)
	res, err := uc.ReadBy(outletId, productId)

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler ProductPriceHandler) Add(ctx *fiber.Ctx) (err error) {

	// Parse & Checking input
	input := new(request.ProductPriceRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := usecase.NewProductPriceUseCase(handler.UcContract)
	res, err := uc.Add(input)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler ProductPriceHandler) Edit(ctx *fiber.Ctx) (err error) {
	// Get Param
	ID := ctx.Params("id")

	// Parse & Checking input
	input := new(request.ProductPriceRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := usecase.NewProductPriceUseCase(handler.UcContract)
	res, err := uc.Edit(input, ID)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler ProductPriceHandler) DeleteByID(ctx *fiber.Ctx) (err error) {
	// Get Param
	outletId := ctx.Params("outlet_id")
	productId := ctx.Params("product_id")

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := usecase.NewProductPriceUseCase(handler.UcContract)

	if err := uc.DeleteBy(outletId, productId); err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
}
