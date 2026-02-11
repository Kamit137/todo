package sql

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func RegDb(email, password string)error{
	db, err := sql.Open("sqlite3", "db.db")
	if err!= nil{
		log.Fatal("Fail open Db",err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQE NOT NULL,
		password TEXT NOT NULL,
		date_create DATETIME DEFAULT CURRENT_TIMESTAMP);")
	if err != nil{
		log.Fatal("ER create db",err)
		return err
	}
	row, err := db.Query("SELECT email FROM users WHERE email = ?", email)
	if err != nil {
		return err
	}
	defer rows.Close()
	if row == nil{
		db.Exec("INSERT INTO users(email, password) VALUE (?,?)", email, password)

		return nil
	}else{
		return errors.New("emeail exist")
	}



}