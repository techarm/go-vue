package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	var config config
	config.port = 8081

	infoLog := log.New(os.Stdout, "[INFO ] ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	var app application
	app.config = config
	app.infoLog = infoLog
	app.errorLog = errorLog

	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (a application) serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Ok      bool   `json:"ok"`
			Message string `json:"message"`
		}

		payload.Ok = true
		payload.Message = "Hello world"
		res, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			a.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	a.infoLog.Printf("start http server and listening on 127.0.0.1:%d", a.config.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", a.config.port), nil)
}
