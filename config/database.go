package config

import (
	"database/sql"
	"fmt"
)

func DatabaseInit() *sql.DB {
	db, err := sql.Open("mysql", "root:rakamin@tcp(127.0.0.1:3306)/coba")

	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(100) //membatasi open koneksi ke database
	db.SetMaxIdleConns(80)  // membatasi idle koneksi ke database
	fmt.Println("connected to database")

	return db
}
