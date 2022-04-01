package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/techarm/go-vue/vue-api/internal/driver"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *driver.DB
}

func main() {
	var config config
	config.port = 8081

	infoLog := log.New(os.Stdout, "[INFO ] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := "host=localhost port=5432 user=postgres password=password dbname=vueapi sslmode=disable timezone=UTC connect_timeout=5"
	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		errorLog.Fatalln(err)
	}
	infoLog.Println("connect to database.")

	var app application
	app.config = config
	app.infoLog = infoLog
	app.errorLog = errorLog
	app.db = db

	err = app.serve()
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
