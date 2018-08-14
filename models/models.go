package models

import (
	"database/sql"
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection
var SQL *sql.DB

func init() {
	setPop()
	setOther()

}

func setPop() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}

func setOther() {
	var err error
	connStr := envy.Get("DATABASE_URL", "postgres://postgres:postgres@127.0.0.1:5432/certmanager_development?sslmode=disable")
	SQL, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
