package db

import (
	"database/sql"
	"os"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
	conf "github.com/minhajuddinkhan/gopansy/config"
	"github.com/minhajuddinkhan/gopansy/constants"
	"github.com/tkanos/gonfig"
)

//Migrate Migrate
func Migrate() {

	var configuration conf.Configuration
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}

	path := "../../config/" + conf.GetEnvPath(env)
	migrationPath := "../../db/migrations"

	err := gonfig.GetConf(path, &configuration)
	HandleBootstrapError(err)
	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	HandleBootstrapError(err)

	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, migrationPath)
	err = migrator.Migrate()
	HandleBootstrapError(err)

}
