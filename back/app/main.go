package main

import (
	"fmt"

	"./gormdb"
	"./query"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ほしそうなクエリ
// - POST GPSからレコードを受け取る
// - GET  レコードを条件づけて取得 例えば直近5分
// - GET  レコード全件取得

func main() {

	gormdb.CreateDB()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}}))
	r.GET("/ahchoo/get", query.GetOneRecord)
	r.GET("/ahchoo/getall", query.GetAllRecord)
	r.GET("/ahchoo/one-hour", query.GetOneHour)
	r.GET("/ahchoo/one-day", query.GetOneDay)
	r.GET("/ahchoo/seven-days", query.GetSevenDays)
	r.GET("/ahchoo/thiry-days", query.GetThirtyDays)
	r.POST("/ahchoo/post", query.PostFromApp)
	r.Run(":8080")
	fmt.Println("after run")

}
