package db

import (
	"log"

	conf "github.com/minhajuddinkhan/gopansy/config"
)

var configuration conf.Configuration

//Migrate Migrate
func Migrate() {

	// env := os.Getenv("ENV")
	// if len(env) == 0 {
	// 	env = "dev"
	// }
	// err := gonfig.GetConf("./config/"+conf.GetEnvPath(env), &configuration)
	// handleBootstrapError(err)

	// conf.SetConfig(configuration)

	// db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	// handleBootstrapError(err)

	// migrator, _ :=
	// err = migrator.Migrate()

	// return err

}

func handleBootstrapError(err error) {
	if err != nil {
		log.Fatal("SOMETHING WENT WRONG.", err)
	}
}
