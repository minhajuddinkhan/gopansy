package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"database/sql"

	"github.com/DavidHuie/gomigrate"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	constants "github.com/minhajuddinkhan/gopansy/constants"
	routes "github.com/minhajuddinkhan/gopansy/router"
	"github.com/urfave/negroni"
)

func negroLoggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	db, err := sql.Open("postgres", ConfDev.ConnectionString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx := context.WithValue(r.Context(), constants.DbKey, db)
	r = r.WithContext(ctx)
	next.ServeHTTP(rw, r)

}

func main() {

	db, err := sql.Open(constants.DbType, ConfDev.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./db/migrations")
	err = migrator.Migrate()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := mux.NewRouter()
	mux.HandleFunc("/", routes.SayHello)

	n := negroni.Classic()
	n.UseFunc(negroLoggerMiddleware)
	n.UseHandler(mux)

	svr := http.Server{
		Addr:         ConfDev.Addr,
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	svr.ListenAndServe()

}
