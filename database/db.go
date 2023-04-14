package database

import (
	"database/sql"
	"fmt"
	"time"

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
func TestUser(A string, B string, C string) bool {
	switch B {
	case `email`:
		sqlStatement := `SELECT password FROM app.info WHERE email = $1;`
		var password string
		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		row := star.QueryRow(sqlStatement, A)
		row.Scan(&password)
		if password == C {
			return true
		} else {
			return false
		}
	case `username`:
		sqlStatement := `SELECT  password FROM app.info WHERE username = $1;`
		var password string
		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		row := star.QueryRow(sqlStatement, A)
		row.Scan(&password)
		if password == C {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}
func SearchUSER(A string, B string) int {
	switch B {
	case `email`:
		sqlStatement := `SELECT id FROM app.info WHERE email = $1;`
		var id int
		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		row := star.QueryRow(sqlStatement, A)
		switch err := row.Scan(&id); err {
		case sql.ErrNoRows:
			return -2
		case nil:
			return id
		default:
			return -1
		}
	case `username`:
		sqlStatement := `SELECT id FROM app.info WHERE username = $1;`
		var id int
		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		row := star.QueryRow(sqlStatement, A)
		switch err := row.Scan(&id); err {
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
func RemoveUSER(A string) int {
	var id int = 0
	sqlStatement := ` DELETE FROM app.info WHERE id = $1 RETURNING id; `
	err := star.QueryRow(sqlStatement, A).Scan(&id)
	if err != nil {
		id = -1
		return id

	} else {
		return id
	}

}

func UpdateUSER(A string, B string, C string) int {
	switch B {
	case `email`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET email = $1 WHERE id = $2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `username`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET username = $1 WHERE id = $2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `password`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET password = $1 WHERE id = $2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `first_name`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET first_name = $1 WHERE id = $2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `last_name`:
		var id int = 0
		sqlStatement := `UPDATE app.info SET last_name = $1 WHERE id = $2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	default:
		return -1
	}
}
func InsertNote(A string, B string, C int) {
	tme := time.Now().Format("01-02-2006 15:04:05")
	sqlStatement := `
INSERT INTO app.notes (title,description,userid,last_edited)
VALUES ($1, $2, $3, $4) `
	_, err := star.Exec(sqlStatement, A, B, C, tme)
	if err != nil {
		panic(err)
	}

}
func DeleteNote(A string) int {
	var id int = 0

	sqlStatement := ` DELETE FROM app.notes WHERE id = $1 RETURNING id; `
	err := star.QueryRow(sqlStatement, A).Scan(&id)
	if err != nil {
		id = -1
		return id

	} else {
		return id
	}
}
func UpdateNotes(A string, B string, C string) int {
	switch B {
	case `title`:
		var id int = 0
		tme := time.Now().Format("01-02-2006 15:04:05")
		sqlStatement := `UPDATE app.notes SET title = $1, last_edited = $3  WHERE id = $2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C, tme).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	case `description`:
		var id int = 0
		tme := time.Now().Format("01-02-2006 15:04:05")
		sqlStatement := `UPDATE app.notes SET description = $1, last_edited = $3 WHERE id = $2 RETURNING id;`
		err := star.QueryRow(sqlStatement, A, C, tme).Scan(&id)
		if err != nil {
			panic(err)
		}
		return id
	default:
		return -1
	}
}

type Note struct {
	Title       string
	Description string
	Last_edited string
}

func GetNote(A string) []Note {
	shay := []Note{}
	sqlStatement := `SELECT id, title, description, last_edited FROM app.notes  WHERE userid = $1;`
	rows, err := star.Query(sqlStatement, A)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var description string
		var last_edited string
		err = rows.Scan(&id, &title, &description, &last_edited)
		var strategy Note
		strategy.Title = title
		strategy.Description = description
		strategy.Last_edited = last_edited
		if err != nil {
			// handle this error
			panic(err)
		}
		shay = append(shay, strategy)
	}
	return shay
}
