package modelx

import (
	"database/sql"
	"errors"

	"github.com/rmsj/stock/db/models"
	. "github.com/vattle/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

// UserByEmailAndPassword finds a user by email and password
func UserByEmailAndPassword(db *sql.DB, email string, password string) (user *models.User, err error) {

	user, err = models.Users(db, Where("email = ?", email)).One()
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid user name or password")
	}

	return
}

// UserEmailInUse checks if email is already in use
func UserEmailInUse(db *sql.DB, email string) (exists bool, err error) {
	exists, err = models.Users(db, Where("email = ?", email)).Exists()
	return
}

// UserEmailInUse checks if email is already in use
func UserFromSession(db *sql.DB, sessionID string) (user *models.User, err error) {

	session, err := models.Sessions(db, Where("id = ?", sessionID)).One()
	if err != nil {
		return
	}

	if len(session.ID) > 0 {
		user, err = session.User(db).One()
	}

	return
}
