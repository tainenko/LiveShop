package databases

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)
var DB *sql.DB
func Connect() {
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	connStr := "user=pqgotest dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB=db
	CreateUserTable()
}
