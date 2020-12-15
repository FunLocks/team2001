package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	r.POST("/post", postFromApp("hoge"))
	r.Run()

}

func postFromApp(arg string) gin.HandlerFunc {

	// DBに書き込む処理をする
	return func(c *gin.Context) {
		var loc Location
		if err := c.ShouldBindJSON(&loc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		loc.ID = 0
		loc.CreatedAt, _ = time.Parse("20006-01-02", c.Param("createdat"))
		loc.Latitude, _ = strconv.ParseFloat(c.Param("latitude"), 64)
		loc.Longitude, _ = strconv.ParseFloat(c.Param("longitude"), 64)
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
	// db.NewRecord(loc)
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
	CreatedAt time.Time `json:"createdAt"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}
