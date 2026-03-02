package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/yaredow/goapi/api"
)

var UnauthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
		}

		var logInDetails *tools.LoginDetails
		logInDetails = (*&database).GetUserLoginDetails(username)

		if logInDetails == nil || (token != (*logInDetails).AuthToken) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)

		}

		next.ServeHTTP(w, r)
	})
}
