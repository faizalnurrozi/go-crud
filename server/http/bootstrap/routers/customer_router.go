package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type CustomerRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route CustomerRoute) RegisterRoute() {
	handler := handlers.CustomerHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	customerRoutes := route.RouteGroup.Group("/customers")
	customerRoutes.Use(jwtMiddleware.New)

	customerRoutes.Get("", handler.Browse)
	customerRoutes.Get("/:id", handler.ReadByID)
	customerRoutes.Post("", handler.Add)
	customerRoutes.Put("/:id", handler.Edit)
	customerRoutes.Delete("/:id", handler.DeleteByID)
}
