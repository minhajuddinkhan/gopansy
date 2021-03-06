package main

import (
	"log"
	"os"

	conf "github.com/minhajuddinkhan/gopansy/config"
	db "github.com/minhajuddinkhan/gopansy/db"
	"github.com/minhajuddinkhan/gopansy/helpers"
	"github.com/minhajuddinkhan/gopansy/middlewares"
	"github.com/minhajuddinkhan/gopansy/router"
	"github.com/minhajuddinkhan/gopansy/server"
	"github.com/tkanos/gonfig"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

const (
	migrate  = "migrate"
	seedup   = "seed-up"
	seeddown = "seed-down"
	serve    = "serve"
	start    = "start"
)

func upRouter() *negroni.Negroni {
	n := negroni.Classic()
	n.UseFunc(middlewares.DecodeJWT)
	n.UseFunc(middlewares.SetDbCtx)
	n.UseHandler(router.Initiate())
	return n

}

func upConfig(c *conf.Configuration) {

	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	path := conf.GetEnvPath(env)
	err := gonfig.GetConf("./config/"+path, c)

	if err != nil {
		helpers.HandleBootstrapError(err)
	}

}

func main() {

	app := cli.NewApp()
	app.Name = "Pansy"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {

		var configuration conf.Configuration
		var s server.Server
		upConfig(&configuration)
		dbPath := "./db/migrations"
		conf.SetConfig(configuration)

		switch c.Args().First() {

		case start:
			migrator := db.Migrator{
				Conf: configuration,
			}
			err := migrator.Migrate(dbPath)
			if err != nil {
				log.Fatal(err)
			}

			seeder := db.Seeder{
				Conf: configuration,
			}
			err = seeder.Seed()
			if err != nil {
				log.Fatal(err)
			}
			s = server.Server{
				Conf:   configuration,
				Router: upRouter(),
			}
			return s.Start()

		case migrate:
			migrator := db.Migrator{
				Conf: configuration,
			}
			return migrator.Migrate(dbPath)

		case seedup:
			seeder := db.Seeder{
				Conf: configuration,
			}
			return seeder.Seed()

		case seeddown:
			seeder := db.Seeder{
				Conf: configuration,
			}
			return seeder.UnSeed()

		case serve:
			s = server.Server{
				Conf:   configuration,
				Router: upRouter(),
			}
			return s.Start()

		default:
			s = server.Server{
				Conf:   configuration,
				Router: upRouter(),
			}
			return s.Start()
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
