package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ProductRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ProductRoute) RegisterRoute() {
	handler := handlers.ProductHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	productRoutes := route.RouteGroup.Group("/products")
	productRoutes.Use(jwtMiddleware.New)

	productRoutes.Get("", handler.Browse)
	productRoutes.Get("/:id", handler.ReadByID)
	productRoutes.Post("", handler.Add)
	productRoutes.Put("/:id", handler.Edit)
	productRoutes.Delete("/:id", handler.DeleteByID)
}
