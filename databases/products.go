package databases

func CreateProductTable(){
	DB.Query(`
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