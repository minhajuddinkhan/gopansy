package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	conf "github.com/minhajuddinkhan/gopansy/config"
	middlewares "github.com/minhajuddinkhan/gopansy/middlewares"
	router "github.com/minhajuddinkhan/gopansy/router"
	"github.com/tkanos/gonfig"
	"github.com/urfave/negroni"
)

var configuration conf.Configuration

func main() {

	err := gonfig.GetConf("./config/"+GetEnvPath(), &configuration)
	handleBootstrapError(err)

	conf.SetConfig(configuration)

	n := negroni.Classic()
	n.UseFunc(middlewares.SetDbCtx)
	n.UseHandler(router.Initiate())

	svr := http.Server{
		Addr:         configuration.Addr,
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	svr.ListenAndServe()

}

func handleBootstrapError(err error) {
	if err != nil {
		log.Fatal("SOMETHING WENT WRONG.", err)
	}
}

//GetEnvPath GetEnvPath
func GetEnvPath() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	return conf.GetEnvPath(env)

}
