package query

import (
	"net/http"
	"time"

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
	result := gormdb.Location{}
	db.Model(&gormdb.Location{}).Last(&result)

	defer db2.Close()
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

// GetOneHour 1時間分
func GetOneHour(c *gin.Context) {
	db := gormdb.GormConnect().Model(&gormdb.Location{})
	db2, _ := db.DB()
	result := []gormdb.Location{}

	hour := time.Now().Add(-time.Hour)
	db.Where("created_at > ?", hour).Order("id").Find(&result)
	defer db2.Close()
	c.JSON(http.StatusOK, result)
}

// GetOneDay 1日分
func GetOneDay(c *gin.Context) {
	db := gormdb.GormConnect().Model(&gormdb.Location{})
	db2, _ := db.DB()
	result := []gormdb.Location{}
	today := time.Now().AddDate(0, 0, -1)
	db.Where("created_at > ?", today).Order("id").Find(&result)
	// db.Order("id").Find(&result).Where("created_at BETWEEN 2000-01-01T00:00:00Z AND 2000-01-08T00:00:00Z")
	defer db2.Close()
	c.JSON(http.StatusOK, result)
}

// GetSevenDays 7日分
func GetSevenDays(c *gin.Context) {
	db := gormdb.GormConnect().Model(&gormdb.Location{})
	db2, _ := db.DB()
	result := []gormdb.Location{}
	lastWeek := time.Now().AddDate(0, 0, -7)
	db.Where("created_at > ?", lastWeek).Order("id").Find(&result)
	defer db2.Close()
	c.JSON(http.StatusOK, result)
}

// GetThirtyDays 30日分
func GetThirtyDays(c *gin.Context) {
	db := gormdb.GormConnect().Model(&gormdb.Location{})
	db2, _ := db.DB()
	result := []gormdb.Location{}
	lastMonth := time.Now().AddDate(0, -1, 0)
	db.Where("created_at > ?", lastMonth).Order("id").Find(&result)
	defer db2.Close()
	c.JSON(http.StatusOK, result)
}
