package query

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../geocode"
	"../geoparser"
	"../gormdb"
)

// PostFromApp Andoroidが呼ぶ本体
func PostFromApp() gin.HandlerFunc {

	// DBに書き込む処理をする
	return func(c *gin.Context) {
		var loc gormdb.Location
		if err := c.ShouldBindJSON(&loc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		lat, lon := loc.RawLatitude, loc.RawLongitude
		var geo geoparser.Coord
		geo.Latitude = lat
		geo.Longitude = lon
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
}

// GetOneRecord GETメソッドの本体
func GetOneRecord() gin.HandlerFunc {
	// DBから取得する処理
	db := gormdb.GormConnect()
	result := map[string]interface{}{}
	db.Model(&gormdb.Location{}).Last(&result)
	return func(c *gin.Context) {
		c.JSON(200, result)
	}
}

// GetAllRecord レコードを全県取得
func GetAllRecord() gin.HandlerFunc {
	db := gormdb.GormConnect().Model(&gormdb.Location{})
	result := []map[string]interface{}{}
	db.Order("id").Find(&result)
	return func(c *gin.Context) {
		c.JSON(200, result)
	}
}
