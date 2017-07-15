package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	//Get user's class this semester.
	r.GET("/123", func(c *gin.Context) {
		db, err := gorm.Open("sqlite3", "/class.db")
		defer db.Close()

		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	//Check user's DB to exist.If not exist,initDB().If exist, In user's DB, post data.
	r.POST("/123", func(c *gin.Context) {
		db, err := gorm.Open("sqlite3", "/db/class.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		userId := c.Query("id")
		class := c.Request.class
		absNum := c.Request.absenceNumber

	})
	r.Run()
}

//Creat table and column.
//fnuc initDB() *db {
//1}
