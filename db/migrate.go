package db

import (
	"database/sql"
	"log"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
	conf "github.com/minhajuddinkhan/gopansy/config"
	"github.com/minhajuddinkhan/gopansy/constants"
)

//Migrate Migrate
func Migrate() {

	configuration := conf.GetConfig()
	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	handleBootstrapError(err)

	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./db/migrations")
	err = migrator.Migrate()
	handleBootstrapError(err)

}

//SeederUp SeederUp
func SeederUp() {

	configuration := conf.GetConfig()
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

//SeederDown SeederDown
func SeederDown() {

	configuration := conf.GetConfig()
	db, err := sql.Open(constants.DbType, configuration.ConnectionString)
	handleBootstrapError(err)

	defer db.Close()

	_, err = db.Exec("DELETE from users where id = 1")
	handleBootstrapError(err)

	_, err = db.Exec("DELETE from roles where id = 1")
	handleBootstrapError(err)

}
func handleBootstrapError(err error) {
	if err != nil {
		log.Fatal("SOMETHING WENT WRONG.", err)
	}
}
