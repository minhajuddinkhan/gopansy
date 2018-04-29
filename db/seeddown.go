package db

import (
	"database/sql"
	"os"

	conf "github.com/minhajuddinkhan/gopansy/config"
	"github.com/minhajuddinkhan/gopansy/constants"
	"github.com/tkanos/gonfig"
)

//SeederDown SeederDown
func SeederDown() {

	var configuration conf.Configuration
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}

	path := "../../config/" + conf.GetEnvPath(env)

	err := gonfig.GetConf(path, &configuration)
	handleBootstrapError(err)

	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	handleBootstrapError(err)

	defer db.Close()

	_, err = db.Exec("DELETE from users where id = 1")
	handleBootstrapError(err)

	_, err = db.Exec("DELETE from roles where id = 1")
	handleBootstrapError(err)

}
