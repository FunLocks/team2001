package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ほしそうなクエリ
// - POST GPSからレコードを受け取る
// - GET  レコードを条件づけて取得 例えば直近5分
// - GET  レコード全件取得

func main() {

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

		layout := "2006-01-02 15:04:05"
		loc.ID = 0
		loc.PushTime, _ = time.Parse(layout, c.Param("push_time"))
		loc.Latitude = c.Param("latitude")
		loc.Longitude = c.Param("longitude")
		insertOneRecord(loc)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

// Webが呼ぶやつ
func getFromWeb() gin.HandlerFunc {
	// DBから取得する処理
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "yowa"
	PASS := "yowayowa01"
	PROTOCOL := "tcp(localhost)"
	DBNAME := "test"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("db connectd: ", &db)
	return db
}

func insertOneRecord(loc Location) {
	db := gormConnect()
	defer db.Close()
	db.AutoMigrate(&Location{})
	db.NewRecord(loc)
	db.Create(&loc)

}

func insertMenyRecord(locs []Location) {
	db := gormConnect()
	defer db.Close()
	for _, loc := range locs {
		db.Create(&loc)
	}
}

// Location GPSモジュールから飛んでくるやつ
type Location struct {
	ID        int       `json:"id"`
	PushTime  time.Time `json:"push_time" time_format:"2006-01-02 15:04:05"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
}
