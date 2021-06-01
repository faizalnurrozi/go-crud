package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type SupplierRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route SupplierRoute) RegisterRoute() {
	handler := handlers.SupplierHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	supplierRoutes := route.RouteGroup.Group("/suppliers")
	supplierRoutes.Use(jwtMiddleware.New)

	supplierRoutes.Get("", handler.Browse)
	supplierRoutes.Get("/:id", handler.ReadByID)
	supplierRoutes.Post("", handler.Add)
	supplierRoutes.Put("/:id", handler.Edit)
	supplierRoutes.Delete("/:id", handler.DeleteByID)
}
