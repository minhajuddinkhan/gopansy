package middlewares

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	conf "github.com/minhajuddinkhan/gopansy/config"

	"github.com/minhajuddinkhan/gopansy/constants"
)

func SetDbCtx(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	configuration := conf.GetConfig()
	db, err := sql.Open("postgres", configuration.ConnectionString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.WithValue(r.Context(), constants.DbKey, db)
	r = r.WithContext(ctx)
	next.ServeHTTP(rw, r)

}
