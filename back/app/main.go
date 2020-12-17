package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

// postFromApp Andoroidが呼ぶ本体
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
		insertOneRecord(loc)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

// Webが呼ぶやつ
func getFromWeb() gin.HandlerFunc {
	// DBから取得する処理
	db := gormConnect()
	result := map[string]interface{}{}
	db.Model(&Location{}).Last(&result)
	return func(c *gin.Context) {
		c.JSON(200, result)
	}
}

func gormConnect() *gorm.DB {
	USER := "yowa"
	PASS := "yowayowa01"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "database"
	// DBNAME := "yowayowa" // 本番
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
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

	db.Migrator().AutoMigrate(&Location{
		ID:        0,
		CreatedAt: time.Time{},
		Latitude:  "",
		Longitude: "",
	})
	fmt.Print("test: ")
	fmt.Printf("%v+", &loc)
	db.Create(&loc)

}

func insertMenyRecord(locs []Location) {
	db := gormConnect()
	for _, loc := range locs {
		db.Create(&loc)
	}
}

// Location GPSモジュールから飛んでくるやつ
type Location struct {
	ID          int `gorm:"primary_key"`
	CreatedAt   time.Time
	Latitude    string `json:"latitude" gorm:"size:255"`
	Longitude   string `json:"longitude" gorm:"size:255"`
	Temprature  string `json:"temprature"`
	AirPressure string `json:"AirPressure"`
}
