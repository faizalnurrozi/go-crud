package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

type AdminRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route AdminRoute) RegisterRoute() {
	handler := handlers.AdminHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	adminRoutes := route.RouteGroup.Group("/admin")
	adminRoutes.Use(jwtMiddleware.New)

	adminRoutes.Get("", handler.Browse)
	adminRoutes.Get("/:id", handler.ReadByID)
	adminRoutes.Post("", handler.Add)
	adminRoutes.Put("/:id", handler.Edit)
	adminRoutes.Delete("/:id", handler.DeleteByID)
}
