package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/techarm/go-vue/vue-api/internal/data"
	"github.com/techarm/go-vue/vue-api/internal/driver"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   data.Models
}

func main() {
	var config config
	config.port = 8081

	infoLog := log.New(os.Stdout, "[INFO ] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := driver.ConnectPostgres(os.Getenv("DSN"))
	if err != nil {
		errorLog.Fatalln(err)
	}
	infoLog.Println("connect to database.")

	var app application
	app.config = config
	app.infoLog = infoLog
	app.errorLog = errorLog
	app.models = data.New(db.SQL)

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
	}
}

// serve start the http web server
func (app application) serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	app.infoLog.Printf("start http server and listening on 127.0.0.1:%d", app.config.port)
	return srv.ListenAndServe()
}
