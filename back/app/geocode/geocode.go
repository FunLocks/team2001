package geocode

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"../types"
)

// GetAddressFromCoord 座標から都市を含んだ構造体
func GetAddressFromCoord(latitude, longitude string) *types.AutoGenerated {
	LINK := "https://maps.googleapis.com/maps/api/geocode/json?latlng="
	QUERY := "&location_type=ROOFTOP&result_type=street_address&key="

	var result *types.AutoGenerated

	f, err := os.Open("./apikey/api.txt")
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer f.Close()
	APIKEY := fileToString(f)

	URL := LINK + latitude + "," + longitude + QUERY + APIKEY

	result, err = toJSON(URL)
	if err != nil {
		fmt.Println(err)
		return result
	}
	return result
}

// GetCoordFromAddress 都市から座標を含んだ構造体
func GetCoordFromAddress(city, town string) *types.AutoGenerated {
	LINK := "https://maps.googleapis.com/maps/api/geocode/json?address="
	QUERY := "&language=en&key="

	var result *types.AutoGenerated

	f, err := os.Open("./apikey/api.txt")
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer f.Close()
	APIKEY := fileToString(f)

	// Sumida City みたいに入ってくるとリンクがちぎれてバグる
	URL := LINK + town + "%" + city + QUERY + APIKEY

	result, err = toJSON(URL)
	if err != nil {
		fmt.Println(err)
		return result
	}
	return result
}

func fileToString(fp *os.File) string {
	scanner := bufio.NewScanner(fp)
	var result string
	for scanner.Scan() {
		if x := scanner.Text(); x != "\n" {
			result += x
		}

	}
	if err := scanner.Err(); err != nil {
		// panic(err)
		return "error"
	}
	// バイト型スライスを文字列型に変換してファイルの内容を返却
	return result
}

func toJSON(url string) (*types.AutoGenerated, error) {
	resp, err := http.Get(url)

	// var result types.AutoGenerated
	var data *types.AutoGenerated

	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)

	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		fmt.Println(data.Results)
		fmt.Println(data)
		return data, err
	}
	return data, err
}
