package handlers

import (
	"github.com/faizalnurrozi/go-crud/domain/request"
	"github.com/faizalnurrozi/go-crud/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthenticationHandler struct {
	Handler
}

func (handler AuthenticationHandler) Login(ctx *fiber.Ctx) error {
	input := new(request.LoginRequest)

	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}

	uc := usecase.NewAuthenticationUseCase(handler.UcContract)
	res, err := uc.Login(input)

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

func (handler AuthenticationHandler) GetCurrentUser(ctx *fiber.Ctx) (err error) {
	uc := usecase.NewAuthenticationUseCase(handler.UcContract)
	res, err := uc.GetCurrentUser()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}
