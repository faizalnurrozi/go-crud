package main

import (
	"fmt"
	"github.com/faizalnurrozi/go-crud/domain"
	"github.com/faizalnurrozi/go-crud/server/http/bootstrap"
	"github.com/faizalnurrozi/go-crud/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/xid"
	"log"
	"os"
	"time"
)

var (
	logFormat = `{"host":"${host}","pid":"${pid}","time":"${time}","request-id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","in":"${bytesReceived}","out":"${bytesSent}"}`
)

func main() {
	config, err := domain.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	defer config.DB.Close()

	//init validator
	domain.ValidatorInit()

	app := fiber.New()

	//use case contract init
	ucContract := usecase.Contract{
		ReqID:         xid.New().String(),
		App:           app,
		DB:            config.DB,
		TX:            nil,
		Validate:      domain.ValidatorDriver,
		Translator:    domain.Translator,
		JwtCredential: config.JwtCredential,
		JweCredential: config.JweCredential,
	}

	//bootstrap init
	boot := bootstrap.Bootstrap{
		App:        app,
		Db:         config.DB,
		UcContract: ucContract,
		Validator:  domain.ValidatorDriver,
		Translator: domain.Translator,
	}

	boot.App.Use(recover.New())
	boot.App.Use(requestid.New())
	boot.App.Use(cors.New())
	boot.App.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	boot.RegisterRoute()
	log.Fatal(boot.App.Listen(os.Getenv("APP_HOST")))
}
