package main

import (
	"errors"
	"net/http"
)

func (app *application) AuthoTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := app.models.Token.AuthenticateToken(r)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, errors.New("認証失敗"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
