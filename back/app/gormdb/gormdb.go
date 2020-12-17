package gormdb

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormConnect DBに接続
func GormConnect() *gorm.DB {
	USER := "yowa"
	PASS := "yowayowa01"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "ahchoo"
	// DBNAME := "yowayowa" // 本番
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return db
}

// CreateDB データベースがなかったら自動で作成．
func CreateDB() {
	db := GormConnect()
	if !db.Migrator().HasTable("locations") {
		db.Migrator().CreateTable(&Location{})
	}
}

// InsertOneRecord レコードを1つ挿入
func InsertOneRecord(loc Location) {
	db := GormConnect()

	db.Migrator().AutoMigrate(&Location{
		ID:        0,
		CreatedAt: time.Time{},
		Latitude:  "",
		Longitude: "",
	})
	// fmt.Print("test: ")
	// fmt.Printf("%v+", &loc)
	db.Create(&loc)

}

func insertMenyRecord(locs []Location) {
	db := GormConnect()
	for _, loc := range locs {
		db.Create(&loc)
	}
}

// Location GPSモジュールから飛んでくるやつ
type Location struct {
	ID           int `gorm:"primary_key"`
	CreatedAt    time.Time
	RawLatitude  string `json:"latitude" gorm:"size:255"`
	RawLongitude string `json:"longitude" gorm:"size:255"`
	Latitude     string `gorm:"size:255"`
	Longitude    string `gorm:"size:255"`
	Town         string `json:"town" gorm:"size:255"`
	Temprature   string `json:"temprature"`
	AirPressure  string `json:"AirPressure"`
}
