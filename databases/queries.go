package databases

const (
	IsUserExist = `SELECT id FROM users WHERE email=$1;`
	CreateUserQuery = `INSERT INTO users(id,name,password,email) VALUES (DEFAULT, $1 , $2, $3);`
	QueryUserWithParam = `SELECT id FROM users WHERE email=$1 and password=$2`
)
