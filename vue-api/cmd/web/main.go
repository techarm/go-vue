package main

import (
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

func (app application) serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	app.infoLog.Printf("start http server and listening on 127.0.0.1:%d", app.config.port)
	return srv.ListenAndServe()
}
