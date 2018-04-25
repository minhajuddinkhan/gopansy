package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
	conf "github.com/minhajuddinkhan/gopansy/config"
	"github.com/minhajuddinkhan/gopansy/constants"

	migrate "github.com/minhajuddinkhan/gopansy/db"
	middlewares "github.com/minhajuddinkhan/gopansy/middlewares"
	router "github.com/minhajuddinkhan/gopansy/router"
	"github.com/tkanos/gonfig"
	"github.com/urfave/negroni"
)

var configuration conf.Configuration

func main() {

	bootstrapConfig()
	migrate.Migrate()
	migrate.SeederUp()

	svr := http.Server{
		Addr:         configuration.Addr,
		Handler:      bootstrapRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	svr.ListenAndServe()

}

func bootstrapConfig() {

	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	path := conf.GetEnvPath(env)
	err := gonfig.GetConf("./config/"+path, &configuration)
	handleBootstrapError(err)
	conf.SetConfig(configuration)

}

func bootstrapMigrations() {
	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	handleBootstrapError(err)

	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./db/migrations")
	err = migrator.Migrate()
	handleBootstrapError(err)

	defer db.Close()

}

func bootstrapRouter() *negroni.Negroni {

	n := negroni.Classic()
	n.UseFunc(middlewares.EncodeJWT)
	n.UseFunc(middlewares.SetDbCtx)
	n.UseHandler(router.Initiate())
	return n

}

func handleBootstrapError(err error) {
	if err != nil {
		log.Fatal("SOMETHING WENT WRONG.", err)
		panic(err)
	}
}
