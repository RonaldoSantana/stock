package api

import (
	"time"

	"github.com/rmsj/stock/db/models"
	"github.com/satori/go.uuid"
	"gopkg.in/nullbio/null.v6"
)

// sessionTokenResponse type
type sessionTokenResponse struct {
	Token string `json:"sessionToken"`
}

// create new session on DB
func (api *API) startSession(user *models.User) (sessionID string, err error) {

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

	err = dbSession.Insert(api.DB)
	if err != nil {
		return
	}
	sessionID = dbSession.ID

	return
}
