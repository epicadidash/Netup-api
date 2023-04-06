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

	go func() {
		var id int = 0
		sqlStatement := `
INSERT INTO app.info ( email, first_name, last_name, username, password)
VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err := star.QueryRow(sqlStatement, A, B, C, D, E).Scan(&id)
		if err != nil {
			panic(err)
		}

	}()
}
func SearchUSER(A string, B string) int {
	switch B {
	case `email`:
		sqlStatement := `SELECT id, email FROM app.info WHERE email = $1;`
		var email string
		var id int
		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		row := star.QueryRow(sqlStatement, A)
		switch err := row.Scan(&id, &email); err {
		case sql.ErrNoRows:
			return -1
		case nil:
			return id
		default:
			return -1
		}
	case `username`:
		sqlStatement := `SELECT id, email FROM app.info WHERE username = $1;`
		var email string
		var id int
		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		row := star.QueryRow(sqlStatement, A)
		switch err := row.Scan(&id, &email); err {
		case sql.ErrNoRows:
			return -1
		case nil:
			return id
		default:
			return -1
		}
	default:
		return -1
	}
}
func RemoveUSER(A int) {
	var id int = 0
	sqlStatement := `DELETE FROM app.info WHERE id = $1 RETURNING id;`
	err := star.QueryRow(sqlStatement, A).Scan(&id)
	if err != nil {
		panic(err)
	}
}

func UpdateUSER(A string, B string, C int) int {
	switch B {
	case `email`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET email = $1 WHERE id =$2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `username`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET username = $1 WHERE id =$2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `password`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET password = $1 WHERE id =$2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `first_name`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET first_name = $1 WHERE id =$2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `last_name`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET last_name = $1 WHERE id =$2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	default:
		return -1
	}
}
