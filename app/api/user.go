package api

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/helper"
	"github.com/rmsj/stock/app/security"
	"github.com/rmsj/stock/db/models"
	"github.com/rmsj/stock/db/modelx"
	"github.com/satori/go.uuid"
)

// type to help with login
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterUser registers a user in DB - an admin user
func (api *API) RegisterUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	user := LoginUser{}
	err := helper.LoadFromJSON(request.Body, &user)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("User: %t, %s", user, user)

	exists, err := modelx.UserEmailInUse(api.DB, user.Email)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	if exists {
		api.WriteError(writer, http.StatusBadRequest, "email already registered")
		return
	}

	password, err := security.HashPassword(user.Password)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	dbUser := &models.User{
		ID:       uuid.NewV4().String(),
		Email:    user.Email,
		Password: password,
		Status:   "active",
	}
	err = dbUser.Insert(api.DB)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	api.registerSessionAndToken(writer, dbUser)
}

func (api *API) Login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	user := LoginUser{}
	err := helper.LoadFromJSON(request.Body, &user)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Read user")

	dbUser, err := modelx.UserByEmailAndPassword(api.DB, user.Email, user.Password)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	// do we have a user
	if len(dbUser.ID) == 0 {
		api.WriteError(writer, http.StatusBadRequest, "Invalid user name or password")
		return
	}

	api.registerSessionAndToken(writer, dbUser)
}

// common actions to create session and token
func (api *API) registerSessionAndToken(writer http.ResponseWriter, user *models.User) {
	sessionID, err := api.startSession(user)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	tokenString, err := security.GenerateToken(sessionID)
	if err != nil {
		api.WriteError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	err = api.WriteJSON(writer, http.StatusOK, sessionTokenResponse{Token: tokenString})
	if err != nil {
		api.WriteError(writer, http.StatusInternalServerError, err.Error())
	}
}

// User retrieves one user
func (api *API) GetUser(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// the middleware should have set this now
	token := req.Context().Value("user")

	claims := token.(*jwt.Token).Claims.(*jwt.StandardClaims)

	user, err := modelx.UserFromSession(api.DB, claims.Id)
	if err != nil {
		api.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	api.WriteJSON(writer, http.StatusOK, user)
}
