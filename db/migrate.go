package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/DavidHuie/gomigrate"
	conf "github.com/minhajuddinkhan/gopansy/config"
	"github.com/minhajuddinkhan/gopansy/constants"
	"github.com/tkanos/gonfig"
)

var configuration conf.Configuration

func main() {

	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	err := gonfig.GetConf("./config/"+conf.GetEnvPath(env), &configuration)
	handleBootstrapError(err)

	conf.SetConfig(configuration)

	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	handleBootstrapError(err)

	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./db/migrations")
	err = migrator.Migrate()

}

func handleBootstrapError(err error) {
	if err != nil {
		log.Fatal("SOMETHING WENT WRONG.", err)
	}
}
