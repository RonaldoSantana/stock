package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/controllers"
)

// Route type
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc httprouter.Handle
	Middleware  []Middleware
}

func (app *Application) getRoutes() []Route {

	userCtl := controllers.NewUserController()
	productCtl := controllers.NewProductController()

	apiRoutes := []Route{

		// signing and sign out route
		Route{"Register", "POST", "/register", userCtl.Register, APINoAuthMiddleware},
		Route{"Login", "POST", "/login", userCtl.Login, APINoAuthMiddleware},
		// User routes
		Route{"GetUser", "GET", "/user", userCtl.Get, APIAuthMiddleware},
		// Product related routes
		Route{"AddProduct", "POST", "/product", productCtl.Get, APIAuthMiddleware},
	}

	return createRoutes(apiRoutes)
}

// creates all routes based on slice of Routes
func createRoutes(route ...[]Route) (routes []Route) {
	for _, r := range route {
		routes = append(routes, r...)
	}
	return
}
