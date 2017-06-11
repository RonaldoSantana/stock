package controllers

import (
	"fmt"
	"net/http"

	"time"

	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/helper"
	"github.com/rmsj/stock/app/security"
	"github.com/rmsj/stock/db/models"
	"github.com/rmsj/stock/db/modelx"
	"github.com/satori/go.uuid"
	"gopkg.in/nullbio/null.v6"
)

// type to help with login
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserController type
type UserController struct {
	baseController
}

// NewUserController returns a UserController
func NewUserController(db *sql.DB, env map[string]interface{}) *UserController {
	return &UserController{
		baseController{
			Env: env,
			DB:  db,
		},
	}
}

// RegisterUser registers a user in DB - an admin user
func (uc *UserController) RegisterUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	user := LoginUser{}
	err := helper.LoadFromJSON(request.Body, &user)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	exists, err := modelx.UserEmailInUse(uc.DB, user.Email)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	if exists {
		uc.WriteError(writer, http.StatusBadRequest, "email already registered")
		return
	}

	password, err := security.HashPassword(user.Password)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	dbUser := &models.User{
		ID:       uuid.NewV4().String(),
		Email:    user.Email,
		Password: password,
		Status:   "active",
	}
	err = dbUser.Insert(uc.DB)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	uc.registerSessionAndToken(writer, dbUser)
}

// Login - logs user in the system
func (uc *UserController) Login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	user := LoginUser{}
	err := helper.LoadFromJSON(request.Body, &user)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Read user")

	dbUser, err := modelx.UserByEmailAndPassword(uc.DB, user.Email, user.Password)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	// do we have a user
	if len(dbUser.ID) == 0 {
		uc.WriteError(writer, http.StatusBadRequest, "Invalid user name or password")
		return
	}

	uc.registerSessionAndToken(writer, dbUser)
}

// GetUser - User retrieves one user
func (uc *UserController) GetUser(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// the middleware should have set this now
	token := req.Context().Value("user")

	claims := token.(*jwt.Token).Claims.(*jwt.StandardClaims)

	user, err := modelx.UserFromSession(uc.DB, claims.Id)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	uc.WriteJSON(writer, http.StatusOK, user)
}

// common actions to create session and token
func (uc *UserController) registerSessionAndToken(writer http.ResponseWriter, user *models.User) {
	sessionID, err := uc.startSession(user)
	if err != nil {
		uc.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	tokenString, err := security.GenerateToken(sessionID)
	if err != nil {
		uc.WriteError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	err = uc.WriteJSON(writer, http.StatusOK, sessionTokenResponse{Token: tokenString})
	if err != nil {
		uc.WriteError(writer, http.StatusInternalServerError, err.Error())
	}
}

// create new session on DB
func (uc *UserController) startSession(user *models.User) (sessionID string, err error) {

	dbSession := models.Session{
		ID:     uuid.NewV4().String(),
		UserID: user.ID,
		LoginTime: null.Time{
			Valid: true,
			Time:  time.Now(),
		},
		LastSeen: null.Time{
			Valid: true,
			Time:  time.Now(),
		},
	}

	err = dbSession.Insert(uc.DB)
	if err != nil {
		return
	}
	sessionID = dbSession.ID

	return
}
