package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type MerhcantRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route MerhcantRoute) RegisterRoute() {
	handler := handlers.MerchantHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	merchantRoutes := route.RouteGroup.Group("/merchants")
	merchantRoutes.Use(jwtMiddleware.New)

	merchantRoutes.Get("", handler.Browse)
	merchantRoutes.Get("/:id", handler.ReadByID)
	merchantRoutes.Post("", handler.Add)
	merchantRoutes.Put("/:id", handler.Edit)
	merchantRoutes.Delete("/:id", handler.DeleteByID)
}
