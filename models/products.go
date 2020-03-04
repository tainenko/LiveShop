package models

import "github.com/LiveShop/databases"

func CreateProductTable() {
	databases.DB.Query(`
		CREATE TABLE IF NOT EXISTS product(
			goodId INTEGER PRIMARY KEY,
			name VARCHAR(80),
			price INTEGER,
			promoMsg VARCHAR(80),
			image VARCHAR(1024),
			url VARCHAR(1024),
			description VARCHAR(1024),
			cateId INTEGER REFERENCES category (cateId) 
		)`,
	)
}

type Product struct {
	GoodId      int    `json:"goodid"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	PromoMsg    string `json:"promomsg"`
	Image       string `json:"image"`
	Url         string `json:"url"`
	Description string `json:"description"`
	CateId      string `json:"cateid"`
}
