package databases

import (
	"databace/sql"
	_ "github/lib/pq"
	"log"
)
var DB *sql.db
func ConnectDB() {
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	connStr := "user=pqotest dbname=pqotest sslmode=verify-full"
	db, err := sql.Open("postgresql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB=db

}
