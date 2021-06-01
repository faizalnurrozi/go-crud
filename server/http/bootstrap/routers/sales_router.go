package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type SalesRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route SalesRoute) RegisterRoute() {
	handler := handlers.SalesHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	salesRoutes := route.RouteGroup.Group("/sales")
	salesRoutes.Use(jwtMiddleware.New)

	salesRoutes.Get("", handler.Browse)
	salesRoutes.Get("/:id", handler.ReadByID)
	salesRoutes.Post("", handler.Add)
	salesRoutes.Put("/:id", handler.Edit)
	salesRoutes.Delete("/:id", handler.DeleteByID)
}
