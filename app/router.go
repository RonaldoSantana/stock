package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewRouter func
func (app *Application) NewRouter() (router *httprouter.Router) {
	router = httprouter.New()

	router.ServeFiles("/resources/*filepath", http.Dir("./assets/"))

	app.Routes = app.getRoutes()
	for _, route := range app.Routes {
		handler := route.HandlerFunc
		for _, middleware := range route.Middleware {
			handler = middleware(handler)
		}

		router.Handle(route.Method, route.Pattern, handler)
	}
	router.NotFound = http.HandlerFunc(app.API.NotFound)
	return
}

// RouterCORS - Cross Origin Resource Sharing
type RouterCORS struct {
	router *httprouter.Router
}

func (rc *RouterCORS) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if origin := request.Header.Get("Origin"); origin != "" {
		writer.Header().Set("Access-Control-Allow-Origin", origin)
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		writer.Header().Set("Server", "Go")
	}
	if request.Method == "OPTIONS" {
		return
	}
	rc.router.ServeHTTP(writer, request)
}
