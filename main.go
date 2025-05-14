package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	// "golang.org/x/crypto/bcrypt"
)

type user struct {
	id          int
	user_name   string
	password    string
	application string
}

const db_name string = "./test.db"

// func hash_passwords(pass []byte) string {
// 	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
// 	if err != nil {
// 		fmt.Println("Could not has password!")
// 	}

// 	return string(hash)
// }

func initilize_db() (string, error) {
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		return "Could not open database", err
	}

	defer db.Close()

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users_apps (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        user_name TEXT,
		password TEXT,
		application TEXT
    );

	CREATE TABLE IF NOT EXISTS user_login (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        user_name TEXT,
		password TEXT,
    );
    `

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return "Could not exectet sql statement", err
	}

	return "Succsesfully Initilized Database!", err
}

func instert_user_data(data []string) (string, error) {
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		return "Could not open database", err
	}

	defer db.Close()

	sqlStmt := `
    INSERT INTO users(user_name, password, application)
	VALUES(?, ?, ?);
    `

	// _, err = db.Exec(sqlStmt, data[0], hash_passwords([]byte(data[1])), data[2])
	_, err = db.Exec(sqlStmt, data[0], data[1], data[2])
	if err != nil {
		return "Could not exectet sql statement", err
	}

	return "Succsesfully Inserted Data Into Database!", err
}

func fetch_all_data() (string, error) {
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		return "Could not open database", err
	}

	defer db.Close()

	sqlStmt := "SELECT * FROM users"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return "Could not fetch data", err
	}

	defer rows.Close()

	for rows.Next() {
		user := user{}

		err := rows.Scan(&user.id, &user.user_name, &user.password, &user.application)
		if err != nil {
			return "Could not scan data", err
		}

		fmt.Printf("User: %d, %s, %s, %s\n", user.id, user.user_name, user.password, user.application)
	}

	if err = rows.Err(); err != nil {
		return "Error from database", err
	}

	return "Succsesfully Fetched Data From Database!", err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Invlaed Commands")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "--init-db":
		data, err := initilize_db()
		if err != nil {
			fmt.Printf(data+": %v", err)
		}

		fmt.Println(data)
	case "--insert-data":
		if len(args) > 0 {
			data, err := instert_user_data(args)
			if err != nil {
				fmt.Printf(data+": %v", err)
			}

			fmt.Println(data)
		}
	case "--insert-user":
		
	case "--fetch":
		data, err := fetch_all_data()
		if err != nil {
			fmt.Printf(data+": %v", err)
		}

		fmt.Println(data)
	}

}
