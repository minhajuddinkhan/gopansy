package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/DavidHuie/gomigrate"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	conf "github.com/minhajuddinkhan/gopansy/config"
	constants "github.com/minhajuddinkhan/gopansy/constants"
	routes "github.com/minhajuddinkhan/gopansy/router"
	"github.com/tkanos/gonfig"
	"github.com/urfave/negroni"
)

var configuration conf.Configuration
var envPath string

func negroLoggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	db, err := sql.Open("postgres", configuration.ConnectionString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.WithValue(r.Context(), constants.DbKey, db)
	r = r.WithContext(ctx)
	next.ServeHTTP(rw, r)

}

func main() {

	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	environments := conf.GetEnvs()
	envPath = environments[env]
	fmt.Println(envPath)
	err := gonfig.GetConf("./config/"+env, &configuration)

	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
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
		Addr:         configuration.Addr,
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	svr.ListenAndServe()

}
