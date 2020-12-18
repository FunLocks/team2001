package main

import (
	"github.com/gin-gonic/gin"

	"./gormdb"
	"./query"
)

// ほしそうなクエリ
// - POST GPSからレコードを受け取る
// - GET  レコードを条件づけて取得 例えば直近5分
// - GET  レコード全件取得

func main() {

	gormdb.CreateDB()

	r := gin.Default()
	r.GET("/get", query.GetOneRecord())
	r.GET("/getall", query.GetAllRecord())
	r.POST("/post", query.PostFromApp())
	r.Run()

}
