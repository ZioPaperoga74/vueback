package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"vue-api/internal/driver"
	"vue-api/internal/data"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger

	models data.Models
}

func main() {
	var cfg config
	cfg.port = 8081

	infoLog := log.New(os.Stdout, "INFO/t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR/t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn:= os.Getenv("DSN")

	db,err:= driver.ConnectPostgres(dsn)
	if err!= nil{
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,

		
		errorLog: errorLog,
		models: data.New(db.SQL),

	
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {

	app.infoLog.Println("Api listening on port", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
