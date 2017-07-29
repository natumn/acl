package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type class struct {
	class  string `json:class`
	day    string `json:day`
	period string `json:period`
	count  int    `json:count`
}

func main() {
	r := gin.Default()

	//Get user's class this semester.
	r.GET("/123", func(c *gin.Context) {
		db := dbConnect()
		defer db.Close()

		db.CreateTable(&semester{})
		db.Model(&semester{}).ModifyColumn()

		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	//Check user's DB to exist.If not exist,initDB().If exist, In user's DB, post data.
	r.POST("/", func(c *gin.Context) {
		db := dbConnect()
		defer db.Close()
		//userId := c.Query("id")
		class := c.PostForm("class")
		absNum := c.PostForm("absNum")

		fmt.Println("%s ,%s ,%s", userId, absNum, class)
	})
	r.Run()
}

//Creat table and column.
//fnuc initDB() *db {
//}

//Connecting to a sqlite3
func dbConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/db/class.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
