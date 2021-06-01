package bootstrap

import (
	"github.com/faizalnurrozi/go-crud/server/http/bootstrap/routers"
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func (boot Bootstrap) RegisterRoute() {
	handlerType := handlers.Handler{
		App:        boot.App,
		UcContract: &boot.UcContract,
		DB:         boot.Db,
		Validate:   boot.Validator,
		Translator: boot.Translator,
	}

	//route for check health
	rootParentGroup := boot.App.Group("/main")
	rootParentGroup.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("auth is working")
	})

	//grouping v1 api
	apiV1 := rootParentGroup.Group("/v1")

	//authentication route
	authenticationRoutes := routers.AuthenticationRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	authenticationRoutes.RegisterRoute()

	//admin route
	adminRoute := routers.AdminRoute{RouteGroup: apiV1, Handler: handlerType}
	adminRoute.RegisterRoute()

	//customer route
	customerRoute := routers.CustomerRoute{RouteGroup: apiV1, Handler: handlerType}
	customerRoute.RegisterRoute()

	//merchant route
	merchantRoute := routers.MerhcantRoute{RouteGroup: apiV1, Handler: handlerType}
	merchantRoute.RegisterRoute()

	//outlet route
	outletRoute := routers.OutletRoute{RouteGroup: apiV1, Handler: handlerType}
	outletRoute.RegisterRoute()

	//product route
	productRoute := routers.ProductRoute{RouteGroup: apiV1, Handler: handlerType}
	productRoute.RegisterRoute()

	//product price route
	productPriceRoute := routers.ProductPriceRoute{RouteGroup: apiV1, Handler: handlerType}
	productPriceRoute.RegisterRoute()

	//supplier route
	supplierRoute := routers.SupplierRoute{RouteGroup: apiV1, Handler: handlerType}
	supplierRoute.RegisterRoute()

	//sales route
	salesRoute := routers.SalesRoute{RouteGroup: apiV1, Handler: handlerType}
	salesRoute.RegisterRoute()
}
