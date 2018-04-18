package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mattes/migrate"

	"github.com/mattes/migrate/database/postgres"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/github"
	"github.com/urfave/negroni"
)

type contextKey string

func (c contextKey) String() string {
	return "mypackage context key " + string(c)
}

var (
	contextKeyAuthtoken = contextKey("auth-token")
	contextKeyAnother   = contextKey("another")
)

func negroLoggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	connStr := "user=pansy-user dbname=pansy-go password=s3cr3tp4ssw0rd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.WithValue(r.Context(), contextKeyAuthtoken, db)
	r = r.WithContext(ctx)
	next.ServeHTTP(rw, r)
}

func main() {

	connStr := "user=pansy-user dbname=pansy-go password=s3cr3tp4ssw0rd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	driver, err := postgres.WithInstance(db)
	m, err := migrate.NewWithDatabaseInstance()

	mux := mux.NewRouter()
	mux.HandleFunc("/", sayHello)
	n := negroni.Classic()
	n.UseFunc(negroLoggerMiddleware)
	n.UseHandler(mux)

	svr := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	svr.ListenAndServe()

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	postgres := r.Context().Value(contextKeyAuthtoken)
	fmt.Println(postgres)
	fmt.Fprintf(w, "Hello World!")
}
