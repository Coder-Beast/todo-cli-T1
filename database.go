package main

import (
	"database/sql"
	"log"

	_"modernc.org/sqlite"
)

var db *sql.DB

func initDB(){
	var err error
	db,err = sql.Open("sqlite","todo.db")
	if err != nil{
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil{
		log.Fatal(err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id TEXT PRIMARY KEY,
        item TEXT,
        completed BOOLEAN
    );`

	_,err = db.Exec(query)

	if err != nil{
		log.Fatal(err)
	}
	log.Println("Database Connected and Table Ready!")

}