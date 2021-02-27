package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joshuaherrera/pern-todo/go_server/go_server/models/psql"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	todos *psql.TodoModel

}

const (
  user = "postgres"
  password = ""
  host = "localhost"
  port = 5432
  database = "perntodo"
)

func main()  {
	addr := flag.String("addr", ":5000", "HTTP Network Address")
	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, database)

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(psqlInfo)
		if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(psqlInfo string) (*sql.DB, error) {
	//https://gowalker.org/github.com/jackc/pgx
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}