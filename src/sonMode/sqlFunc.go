package sonMode

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbPrefix = "root:root@/"
)

//If a mysql database is exist, return a pointer of an open databse,
// or creat a database named by argument.
func SqlDbInitHandle(name string) (db *sql.DB) {
	db, err := sql.Open("mysql", dbPrefix)
	errHandle(err)
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	errHandle(err)
	db.Close()

	db, err = sql.Open("mysql", dbPrefix+name)
	errHandle(err)
	return db
}
