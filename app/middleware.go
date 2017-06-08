package app

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/helpers"
)

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

// Middleware vars
var (
	AppAPI       []Middleware
	AppAPINoAuth []Middleware
)

// Middleware represents one middleware
type Middleware func(h httprouter.Handle) httprouter.Handle

// Middleware func
func (app *Application) Middleware(middleware ...Middleware) []Middleware {
	return middleware
}

func (app *Application) setMiddleware() {
	AppAPI = app.Middleware(app.AuthMiddleware, app.BaseMiddleware)
	AppAPINoAuth = app.Middleware(app.BaseMiddleware)
	return
}

// BaseMiddleware func
func (app *Application) BaseMiddleware(handler httprouter.Handle) httprouter.Handle {
	fn := func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		app.EmailClient = helpers.NewEmailClient(app.Env["PostmarkAPIKey"].(string))
		
		handler(writer, request, params)
		context.Clear(request)
	}
	return httprouter.Handle(fn)
}

// AuthMiddleware func
func (app *Application) AuthMiddleware(handler httprouter.Handle) httprouter.Handle {
	fn := func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		
		// Create the token
		token := jwt.New(jwt.SigningMethodHS256)
		
		// Create a map to store our claims
		claims := token.Claims.(jwt.MapClaims)
	
		// Set token claims
		claims["admin"] = true
		claims["name"] = "Ado Kukic"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		
		// Sign the token with our secret
		tokenString, err := token.SignedString(mySigningKey)
		if err != nil {
			log.Println(err)
			unauthorizedHandler(writer, request)
			return
		}
		
		/* Finally, write the token to the browser window */
		writer.Write([]byte(tokenString))
		handler(writer, request, params)
	}
	return httprouter.Handle(fn)
}


// unauthorizedHandler provides a default HTTP 401 Unauthorized response
func unauthorizedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm=%q`, "Restricted"))
	http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
