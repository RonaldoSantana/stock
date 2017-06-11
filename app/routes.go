package app

import (
	"github.com/julienschmidt/httprouter"
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

	apiRoutes := []Route{

		// signing and sign out route
		Route{"Register", "POST", "/register", app.API.RegisterUser, APINoAuthMiddleware},
		Route{"Login", "POST", "/login", app.API.Login, APINoAuthMiddleware},
		// User routes
		Route{"GetUser", "GET", "/user", app.API.GetUser, APIAuthMiddleware},
		// Product related routes
		Route{"AddProduct", "POST", "/product", app.API.GetUser, APIAuthMiddleware},
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
