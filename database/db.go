package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "netup"
)

var star *sql.DB

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)

	}
	star = db

}
func Ping() {
	err := star.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected")
	}
}
func InsertUSER(A string, B string, C string, D string, E string) {
	var id int = 0
	sqlStatement := `
INSERT INTO app.info ( email, first_name, last_name, username, password)
VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := star.QueryRow(sqlStatement, A, B, C, D, E).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Print(id)
}
func SearchUSER() {
	sqlStatement := `SELECT id, email FROM app.info WHERE id=$1;`
	var email string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := star.QueryRow(sqlStatement, 3)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}
}
