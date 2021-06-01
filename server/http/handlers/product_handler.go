package handlers

import (
	"fmt"
	"github.com/faizalnurrozi/go-crud/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Handler
}

func (handler ProductHandler) Browse(ctx *fiber.Ctx) (err error) {
	// Get Query Param
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	// Database Processing
	uc := usecase.NewProductUseCase(handler.UcContract)
	res, pagination, err := uc.Browse(search, orderBy, sort, page, limit)

	return handler.SendResponse(ctx, ResponseWithMeta, res, pagination, err, http.StatusUnprocessableEntity)
}

func (handler ProductHandler) ReadByID(ctx *fiber.Ctx) (err error) {
	// Get Param
	ID := ctx.Params("id")

	// Database Processing
	uc := usecase.NewProductUseCase(handler.UcContract)
	res, err := uc.ReadBy("p.id", ID, "=")

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler ProductHandler) Add(ctx *fiber.Ctx) (err error) {

	// Upload and save file
	file, err := ctx.FormFile("image")
	if err != nil {
		return err
	}
	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{err.Error()}})
	}

	// Parse & Checking input
	input := make(map[string]interface{})
	input["sku"] = ctx.FormValue("sku")
	input["name"] = ctx.FormValue("name")
	input["image"] = file.Filename
	input["pricePo"] = ctx.FormValue("price_po")
	input["priceSell"] = ctx.FormValue("price_sell")
	input["merchantID"] = ctx.FormValue("merchant_id")

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := usecase.NewProductUseCase(handler.UcContract)
	res, err := uc.Add(input)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler ProductHandler) Edit(ctx *fiber.Ctx) (err error) {

	// Upload and save file
	file, err := ctx.FormFile("image")
	if err != nil {
		return err
	}
	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{err.Error()}})
	}

	// Get Param
	ID := ctx.Params("id")
	input := make(map[string]interface{})
	input["sku"] = ctx.FormValue("sku")
	input["name"] = ctx.FormValue("name")
	input["image"] = ctx.FormValue("image")
	input["pricePo"] = ctx.FormValue("price_po")
	input["priceSell"] = ctx.FormValue("price_sell")
	input["merchantID"] = ctx.FormValue("merchant_id")

	// Database processing
	handler.UcContract.TX, err = handler.DB.Begin()
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := usecase.NewProductUseCase(handler.UcContract)
	res, err := uc.Edit(input, ID)
	if err != nil {
		handler.UcContract.TX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.TX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler ProductHandler) DeleteByID(ctx *fiber.Ctx) (err error) {
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
