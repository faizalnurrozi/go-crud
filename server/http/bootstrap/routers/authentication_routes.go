package routers

import (
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/server/middlewares"
	"github.com/gofiber/fiber/v2"
	//"github.com/faizalnurrozi/go-crud/server/middlewares"
)

type AuthenticationRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route AuthenticationRoutes) RegisterRoute() {
	handler := handlers.AuthenticationHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	authenticationRoutes := route.RouteGroup.Group("/auth")
	authenticationRoutes.Post("/login", handler.Login)
	authenticationRoutes.Use(jwtMiddleware.New).Get("/current", handler.GetCurrentUser)
}
