package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := Open("sqlite3", "class.db")
	defer db.Close()

	db.CreateTable(&Semester{})

}
