package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getdb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	checkerr(err)
	//defer db.Close()
	return db
}
