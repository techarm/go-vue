package main

import (
	"errors"
	"net/http"
)

type credentials struct {
	UserId   string `json:"email"`
	Password string `json:"password"`
}

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data:omitempty"`
}

// Login ログイン処理
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	err := app.readJSON(w, r, &creds)
	if err != nil {
		message := "json format is not corret"
		app.errorJSON(w, errors.New(message))
		app.errorLog.Printf("%s. error: %s", message, err)
		return
	}

	// TODO authenticate
	app.infoLog.Println(creds.UserId, creds.Password)

	// send back a response
	err = app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "success",
	})

	if err != nil {
		app.errorLog.Println(err)
	}
}
