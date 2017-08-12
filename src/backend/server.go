package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type Weekday int

const (
	Monday Weekday = iota
	Tuseday
	Wednesday
	Thursday
	Friday
)

type Class struct {
	Id        int    `json:id`
	className string `json:className`
	count     int    `json:count`
}

type WeekClass struct {
	period int `json:period`
	Class
	day Weekday `json:Weekday`
}

func main() {
	r := gin.Default()

	//Get user's class this semester.
	r.GET("/", func(c *gin.Context) {
		db := dbConnect()
		defer db.Close()

		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	//Check user's DB to exist.If not exist,initDB().If exist, In user's DB, post data.
	r.POST("/", func(c *gin.Context) {
		db := dbConnect()
		defer db.Close()

		w := WeekClass{}

		w.className = c.PostForm("className")
		w.count = strconv.Atoi(c.PostForm("count"))
		w.period = int(c.PostForm("period"))
		w.day = Weekday(c.PostForm("weekday"))

		fmt.Println(w.className, w.day, w.period)
		db.Create(&WeekClass)
		db.Create(&Class)
	})
	r.Run()
}

//Connecting to a sqlite3
func dbConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/class.db")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Class{}, &weekClass{})
	return db
}
