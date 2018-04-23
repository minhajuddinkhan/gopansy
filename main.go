package main

import (
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
	middlewares "github.com/minhajuddinkhan/gopansy/middlewares"
	routes "github.com/minhajuddinkhan/gopansy/router"
	"github.com/tkanos/gonfig"
	"github.com/urfave/negroni"
)

var configuration conf.Configuration
var envPath string

func main() {

	err := gonfig.GetConf("./config/"+GetEnv(), &configuration)

	conf.SetConfig(configuration)
	handleBootstrapError(err)

	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	handleBootstrapError(err)

	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./db/migrations")
	err = migrator.Migrate()
	handleBootstrapError(err)

	defer db.Close()

	mux := mux.NewRouter()
	mux.HandleFunc("/", routes.SayHello)

	n := negroni.Classic()
	n.UseFunc(middlewares.SetDbCtx)
	n.UseHandler(mux)

	svr := http.Server{
		Addr:         configuration.Addr,
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	svr.ListenAndServe()

}

func handleBootstrapError(err) {
	if err != nil {
		log.Fatal("SOMETHING WENT WRONG.", err)
	}
}

func GetEnv() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	return conf.GetEnvPath(env)

}
