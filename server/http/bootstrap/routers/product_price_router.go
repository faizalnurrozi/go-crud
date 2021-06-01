package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ProductPriceRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ProductPriceRoute) RegisterRoute() {
	handler := handlers.ProductPriceHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	productPriceRoutes := route.RouteGroup.Group("/product-prices")
	productPriceRoutes.Use(jwtMiddleware.New)

	productPriceRoutes.Get("/", handler.ReadByID)
	productPriceRoutes.Post("", handler.Add)
	productPriceRoutes.Put("/:id", handler.Edit)
	productPriceRoutes.Delete("/:id", handler.DeleteByID)
}
