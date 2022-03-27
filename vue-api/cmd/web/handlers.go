package main

import (
	"encoding/json"
	"net/http"
)

type credentials struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		message := "json format is not corret"
		writeResponse(app, w, jsonResponse{
			Error:   true,
			Message: message,
		})
		app.errorLog.Printf("%s. error: %s", message, err)
		return
	}

	// TODO authenticate
	app.infoLog.Println(creds.UserId, creds.Password)

	// send back a response
	writeResponse(app, w, jsonResponse{
		Error:   false,
		Message: "success",
	})
}

func writeResponse(app *application, w http.ResponseWriter, payload jsonResponse) {
	out, err := json.Marshal(payload)
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
