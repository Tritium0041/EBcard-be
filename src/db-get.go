package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getdb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/EBcard")
	checkerr(err)
	dbinit(db)
	//defer db.Close()
	return db
}

func dbinit(db *sql.DB) {
	//db := getdb()
	//defer db.Close()
	sql := "CREATE TABLE IF NOT EXISTS INFO(`id` int not null auto_increment,`name` varchar(255) not null,`WXID` varchar(255) not null,`phone` varchar(255) not null,`imageID` varchar(255) not null,`uuid` varchar(255) not null,primary key(`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;"
	_, err := db.Exec(sql)
	checkerr(err)
	sql = "CREATE TABLE IF NOT EXISTS IMAGE(`id` int not null auto_increment,`image` longblob not null,`ImageID` varchar(255) not null,primary key(`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;"
	_, err = db.Exec(sql)
	checkerr(err)
}
