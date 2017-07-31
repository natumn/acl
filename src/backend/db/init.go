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
type Class struct {
  className string `json:class`
  count int `json:count`
  period int `json:period`
}

type class struct {
  gorm.Models
  class Class
	day  Weekday `json:day`
}

func main() {
	db, err := Open("sqlite3", "class.db")
	defer db.Close()

	db.CreateTable(&Semester{})

}
