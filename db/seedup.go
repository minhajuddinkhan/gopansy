package db

import (
	"database/sql"
	"os"

	conf "github.com/minhajuddinkhan/gopansy/config"
	"github.com/minhajuddinkhan/gopansy/constants"
	"github.com/tkanos/gonfig"
)

//SeederUp SeederUp
func SeederUp() {

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

	_, err = db.Exec("INSERT INTO roles (name) SELECT 'admin' WHERE NOT EXISTS ( SELECT id from roles WHERE id = 1)")
	handleBootstrapError(err)
	_, err = db.Exec(`INSERT INTO users
			(username, hashedPassword, roleId, permitOneAllowed, permitTwoAllowed)
			 SELECT 'pancyAdmin', '$2a$05$ZW7dtscHYyl0B7OUlHJ4oOfsJsVt1adavPbpvXi5OjydxM4Tc3QFW', 1, true, true 
			 WHERE NOT EXISTS (SELECT id from users where id = 1)`)
	handleBootstrapError(err)

}
