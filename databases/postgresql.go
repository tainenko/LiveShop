package databases

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)
var DB *sql.DB
func ConnectDB() {
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	connStr := "user=pqotest dbname=pqotest sslmode=verify-full"
	db, err := sql.Open("postgresql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB=db
	CreateUserTable()
}
