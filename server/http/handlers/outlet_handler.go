package handlers

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type OutletHandler struct {
	Handler
}

func (handler OutletHandler) Browse(ctx *fiber.Ctx) (err error) {
	// Get Query Param
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	// Database Processing
	uc := usecase.NewOutletUseCase(handler.UcContract)
	res, pagination, err := uc.Browse(search, orderBy, sort, page, limit)

	return handler.SendResponse(ctx, ResponseWithMeta, res, pagination, err, http.StatusUnprocessableEntity)
}

func (handler OutletHandler) ReadByID(ctx *fiber.Ctx) (err error) {
	// Get Param
	ID := ctx.Params("id")

	// Database Processing
	uc := usecase.NewOutletUseCase(handler.UcContract)
	res, err := uc.ReadBy("o.id", ID, "=")

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler OutletHandler) Add(ctx *fiber.Ctx) (err error) {

	// Parse & Checking input
	input := new(request.OutletRequest)
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
	uc := usecase.NewOutletUseCase(handler.UcContract)
	res, err := uc.Add(input)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler OutletHandler) Edit(ctx *fiber.Ctx) (err error) {
	// Get Param
	ID := ctx.Params("id")

	// Parse & Checking input
	input := new(request.OutletRequest)
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
	uc := usecase.NewOutletUseCase(handler.UcContract)
	res, err := uc.Edit(input, ID)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler OutletHandler) DeleteByID(ctx *fiber.Ctx) (err error) {
	// Get Param
	ID := ctx.Params("id")

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := usecase.NewOutletUseCase(handler.UcContract)

	if err := uc.DeleteBy("id", ID, "="); err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
}
