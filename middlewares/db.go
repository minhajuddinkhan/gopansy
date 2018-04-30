package middlewares

import (
	"context"
	"fmt"
	"net/http"

	boom "github.com/darahayes/go-boom"

	"github.com/jmoiron/sqlx"
	conf "github.com/minhajuddinkhan/gopansy/config"

	"github.com/minhajuddinkhan/gopansy/constants"
)

func SetDbCtx(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	configuration := conf.GetConfig()

	fmt.Println("configuration.ConnectionString", configuration.ConnectionString)
	db, err := sqlx.Open("postgres", configuration.ConnectionString)

	err = db.Ping()
	if err != nil {
		boom.Internal(rw)
		return
	}

	ctx := context.WithValue(r.Context(), constants.DbKey, db)

	r = r.WithContext(ctx)
	next.ServeHTTP(rw, r)

}
