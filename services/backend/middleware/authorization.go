package middleware

import (
	"errors"
	"net/http"

	"github.com/priyanka-choubey/habitrick/backend/api"
	"github.com/priyanka-choubey/habitrick/backend/database"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *database.MySqlDatabase
		database, err = database.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *database.LoginDetails
		loginDetails, err = (*database).GetUserLoginDetails(username)
		if err != nil {
			api.RequestErrorHandler(w, err)
			return
		}

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
