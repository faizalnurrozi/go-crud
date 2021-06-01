package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type OutletRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route OutletRoute) RegisterRoute() {
	handler := handlers.OutletHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	outletRoutes := route.RouteGroup.Group("/outlets")
	outletRoutes.Use(jwtMiddleware.New)

	outletRoutes.Get("", handler.Browse)
	outletRoutes.Get("/:id", handler.ReadByID)
	outletRoutes.Post("", handler.Add)
	outletRoutes.Put("/:id", handler.Edit)
	outletRoutes.Delete("/:id", handler.DeleteByID)
}
