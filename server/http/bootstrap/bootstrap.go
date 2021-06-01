package bootstrap

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/usecase"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Bootstrap struct {
	App        *fiber.App
	Db         *sql.DB
	UcContract usecase.Contract
	Validator  *validator.Validate
	Translator ut.Translator
}
