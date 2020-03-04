package databases

func createAllTable(){
	CreateUserTable()
	CreateProductTable()
}

func CreateProductTable() {
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
func CreateUserTable() {
	DB.Query(`
		CREATE TABLE IF NOT EXISTS users( 
			id serial PRIMARY KEY,
			name VARCHAR (100) NOT NULL,
			password VARCHAR (355) NOT NULL,
			email VARCHAR (355) UNIQUE NOT NULL,
			created_on TIMESTAMP NOT NULL default current_timestamp,
			updated_at TIMESTAMP NOT NULL default current_timestamp 
			)`,
	)
}