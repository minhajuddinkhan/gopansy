package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	conf "github.com/minhajuddinkhan/gopansy/config"

	"github.com/minhajuddinkhan/gopansy/constants"
)

func SetDbCtx(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	configuration := conf.GetConfig()
	db, err := sqlx.Open("postgres", configuration.ConnectionString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.WithValue(r.Context(), constants.DbKey, db)
	r = r.WithContext(ctx)
	next.ServeHTTP(rw, r)

}
