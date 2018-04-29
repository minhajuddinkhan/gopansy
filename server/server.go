package server

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

//Serve Serve
func Serve() {

	bootstrapConfig()

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

func bootstrapRouter() *negroni.Negroni {

	n := negroni.Classic()
	n.UseFunc(middlewares.DecodeJWT)
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
