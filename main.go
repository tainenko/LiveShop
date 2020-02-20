package main

import (
	_ "database/sql"
	"github.com/LiveShop/databases"
	"github.com/LiveShop/router"
	_ "github.com/lib/pq"
)

func init() {
	databases.Connect()
}
func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
