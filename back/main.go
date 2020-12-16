package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ほしそうなクエリ
// - POST GPSからレコードを受け取る
// - GET  レコードを条件づけて取得 例えば直近5分
// - GET  レコード全件取得

func main() {

	createDB()

	r := gin.Default()
	r.GET("/get", getFromWeb())
	r.POST("/post", postFromApp())
	r.Run()

}

func postFromApp() gin.HandlerFunc {

	// DBに書き込む処理をする
	return func(c *gin.Context) {
		var loc Location
		if err := c.ShouldBindJSON(&loc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// layout := "2006-01-02 15:04:05"
		loc.ID = 0
		// loc.PushTime, _ = time.Parse(layout, c.Param("push_time"))
		loc.CreatedAt = time.Now()
		loc.Latitude = c.Param("latitude")
		loc.Longitude = c.Param("longitude")
		fmt.Println(c.Param("latitude"))
		insertOneRecord(loc)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

// Webが呼ぶやつ
func getFromWeb() gin.HandlerFunc {
	// DBから取得する処理
	// var loc Location
	db := gormConnect()
	// result := db.Table("locations").Last(&loc)
	result := map[string]interface{}{}
	db.Model(&Location{}).First(&result)
	return func(c *gin.Context) {
		c.JSON(200, result)
	}
}

func gormConnect() *gorm.DB {
	// DBMS := "mysql"
	USER := "yowa"
	PASS := "yowayowa01"
	PROTOCOL := "tcp(localhost)"
	DBNAME := "test"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	// db, err := gorm.Open(DBMS, CONNECT)
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("db connectd: ", &db)
	return db
}

func createDB() {
	db := gormConnect()
	if !db.Migrator().HasTable("locations") {
		db.Migrator().CreateTable(&Location{})
	}
}

func insertOneRecord(loc Location) {
	db := gormConnect()
	// defer db.Close()

	db.Migrator().AutoMigrate(&Location{
		ID:        0,
		CreatedAt: time.Time{},
		Latitude:  "",
		Longitude: "",
	})
	// db.NewRecord(loc)
	fmt.Print("test: ")
	fmt.Printf("%v+", &loc)
	db.Create(&loc)

}

func insertMenyRecord(locs []Location) {
	db := gormConnect()
	// defer db.Close()
	for _, loc := range locs {
		db.Create(&loc)
	}
}

// Location GPSモジュールから飛んでくるやつ
type Location struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	Latitude  string `json:"latitude" gorm:"size:256"`
	Longitude string `json:"longitude" gorm:"size:256"`
}
