package main

import (
	"log"
	"os"

	m "github.com/minhajuddinkhan/gopansy/db"
	s "github.com/minhajuddinkhan/gopansy/server"
	"github.com/urfave/cli"
)

const (
	migrate  = "migrate"
	seedup   = "seed-up"
	seeddown = "seed-down"
	serve    = "serve"
)

func main() {

	app := cli.NewApp()
	app.Name = "Pansy"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		arg := c.Args().First()

		switch arg {

		case migrate:
			m.Migrate()
			return nil
		case seedup:
			m.SeederUp()
			return nil
		case seeddown:
			m.SeederDown()
			return nil

		case serve:
			s.Serve()
			return nil

		default:
			s.Serve()
			return nil

		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
