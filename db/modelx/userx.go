package modelx

import (
	"database/sql"

	"github.com/rmsj/ecart/app/security"
	"github.com/rmsj/stock/db/models"
	. "github.com/vattle/sqlboiler/queries/qm"
)

// UserByEmailAndPassword finds a user by email and password
func UserByEmailAndPassword(db *sql.DB, email string, password string) (user *models.User, err error) {

	password, err = security.HashPassword(password)
	if err != nil {
		return
	}

	user, err = models.Users(db, Where("email = ? AND password = ?", email, password)).One()

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
