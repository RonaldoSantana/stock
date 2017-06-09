package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/api"
	"github.com/rmsj/stock/app/helper"
	"github.com/rmsj/stock/app/security"
	"github.com/rmsj/stock/db/modelx"
)

// Middleware vars
var (
	APIAuthMiddleware   []Middleware
	APINoAuthMiddleware []Middleware
)

// Middleware represents one middleware
type Middleware func(h httprouter.Handle) httprouter.Handle

// Middleware func
func (app *Application) Middleware(middleware ...Middleware) []Middleware {
	return middleware
}

func (app *Application) setMiddleware() {
	APIAuthMiddleware = app.Middleware(app.AuthMiddleware, app.BaseMiddleware)
	APINoAuthMiddleware = app.Middleware(app.BaseMiddleware)
	return
}

// BaseMiddleware func
func (app *Application) BaseMiddleware(handler httprouter.Handle) httprouter.Handle {
	fn := func(wrt http.ResponseWriter, req *http.Request, params httprouter.Params) {
		app.EmailClient = helper.NewEmailClient(app.Env["PostmarkAPIKey"].(string))

		app.API = api.API{
			Env:         app.Env,
			DB:          app.DB,
			EmailClient: helper.NewEmailClient(app.Env["PostmarkAPIKey"].(string)),
		}

		handler(wrt, req, params)
	}
	return httprouter.Handle(fn)
}

// AuthMiddleware func
func (app *Application) AuthMiddleware(handler httprouter.Handle) httprouter.Handle {

	fn := func(wrt http.ResponseWriter, req *http.Request, params httprouter.Params) {

		extractor := request.HeaderExtractor{
			"Authorization",
		}

		// call back to get the key for verification
		var keyFunc = func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(security.SigningKey), nil
		}
		token, err := request.ParseFromRequestWithClaims(req, extractor, &jwt.StandardClaims{}, keyFunc)

		if err != nil || !token.Valid {
			app.API.WriteError(wrt, http.StatusBadRequest, err.Error())
			return
		}

		claims := token.Claims.(*jwt.StandardClaims)

		_, err = modelx.UserFromSession(app.DB, claims.Id)
		if err != nil {
			unauthorizedHandler(wrt, req)
			return
		}

		newRequest := req.WithContext(context.WithValue(req.Context(), "user", token))

		*req = *newRequest

		handler(wrt, req, params)
	}
	return httprouter.Handle(fn)
}

// unauthorizedHandler provides a default HTTP 401 Unauthorized response
func unauthorizedHandler(wrt http.ResponseWriter, req *http.Request) {
	wrt.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm=%q`, "Restricted"))
	http.Error(wrt, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
