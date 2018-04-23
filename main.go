package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"database/sql"

	"github.com/DavidHuie/gomigrate"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	connStr := "host=db user=pansy-user dbname=pansy-go password=s3cr3tp4ssw0rd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	var str = "INSERT INTO test (id, name) VALUES ('one', 'two')"
	_, err = db.Exec(str)
	if err != nil {
		log.Fatal("CANNOT INSERT.", err)
	}
	ctx := context.WithValue(r.Context(), contextKeyAuthtoken, db)
	r = r.WithContext(ctx)
	next.ServeHTTP(rw, r)

}

func main() {

	connStr := " host=db user=pansy-user dbname=pansy-go password=s3cr3tp4ssw0rd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./db/migrations")
	err = migrator.Migrate()

	if err != nil {
		fmt.Println("migrator error")
		log.Fatal(err)
	}
	defer db.Close()

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
