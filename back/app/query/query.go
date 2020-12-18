package query

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../geocode"
	"../geoparser"
	"../gormdb"
)

// PostFromApp Andoroidが呼ぶ本体
// func PostFromApp(c *gin.Context) gin.HandlerFunc {
func PostFromApp(c *gin.Context) {
	// DBに書き込む処理をする
	// return func() {
	var loc gormdb.Location
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		// return
	}
	lat, lng := loc.RawLatitude, loc.RawLongitude
	var geo geoparser.Coord
	geo.Latitude = lat
	geo.Longitude = lng
	geo.Geodata = *geocode.GetAddressFromCoord(geo.Latitude, geo.Longitude)
	// fmt.Println(geo.GetCityName())
	// fmt.Println(geo.GetTownName())
	// geo.Println()
	var add geoparser.Address
	add.City = geo.GetCityName()
	add.Town = geo.GetTownName()
	add.Geodata = *geocode.GetCoordFromAddress(add.City, add.Town)
	// fmt.Println(add.GetLatitude())
	// fmt.Println(add.GetLongitude())

	loc.Latitude = add.GetLatitude()
	loc.Longitude = add.GetLongitude()
	loc.Town = add.City + " " + add.Town
	// fmt.Printf("%v+", &loc)
	gormdb.InsertOneRecord(loc)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// }

// GetOneRecord GETメソッドの本体
// func GetOneRecord() gin.HandlerFunc {
func GetOneRecord(c *gin.Context) {
	// DBから取得する処理
	db := gormdb.GormConnect()
	db2, _ := db.DB()
	result := map[string]interface{}{}
	db.Model(&gormdb.Location{}).Last(&result)

	defer db2.Close()
	// return func(c *gin.Context) {
	// 	c.JSON(200, result)
	// }
	c.JSON(http.StatusOK, result)
}

// GetAllRecord レコードを全県取得
// func GetAllRecord() gin.HandlerFunc {
func GetAllRecord(c *gin.Context) {
	db := gormdb.GormConnect().Model(&gormdb.Location{})
	db2, _ := db.DB()
	result := []map[string]interface{}{}
	db.Order("id").Find(&result)
	defer db2.Close()
	// return func(c *gin.Context) {
	c.JSON(http.StatusOK, result)
	// }
}
