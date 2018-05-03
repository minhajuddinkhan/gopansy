package router

import (
	"fmt"
	"net/http"

	"github.com/minhajuddinkhan/gopansy/models"

	"github.com/darahayes/go-boom"

	"github.com/jmoiron/sqlx"

	"github.com/minhajuddinkhan/gopansy/helpers"

	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//SayHello say hello dummy func.
func SayHello(w http.ResponseWriter, r *http.Request) {

	db := r.Context().Value(constants.DbKey).(*sqlx.DB)

	tx := db.MustBegin()

	var user models.User
	query := `SELECT * FROM users where id = $1`

	stmt, err := db.Preparex(query)
	defer stmt.Close()

	if err != nil {
		boom.Forbidden(w, err)
		return
	}
	row := stmt.QueryRowx("2")

	err = row.StructScan(&user)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("ERR OCCURED", err)
	}
	//user := r.Context().Value(constants.Authorization)
	helpers.Respond(w, user)
}
