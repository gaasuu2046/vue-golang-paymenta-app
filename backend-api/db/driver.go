package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Conn - sql connection handler
var Conn *sql.DB

// New SQL Halndler - init sql handler
func init() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")

	dbconf := user + ":" + pass + "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error())
	}
	Conn = conn
}
