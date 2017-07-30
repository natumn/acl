package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type Weekday int

count (
  Monday  Weekday = iota
  Tuseday
  Wednesday
  Thursday
  Friday
)

type class struct {
  gorm.Models
  class  string `json:class`
	day  Weekday `json:day`
	period int `json:period`
	count  int    `json:count`
}

func main() {
	db, err := Open("sqlite3", "class.db")
	defer db.Close()

	db.CreateTable(&Semester{})

}
