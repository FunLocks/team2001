package types

// AutoGenerated 自動生成
type AutoGenerated struct {
	Results []Results `json:"results"`
	Status  string    `json:"status"`
}
type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}
type Northeast struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Southwest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Bounds struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Viewport struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}
type Geometry struct {
	Bounds       Bounds   `json:"bounds"`
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}
type Results struct {
	AddressComponents []AddressComponents `json:"address_components"`
	FormattedAddress  string              `json:"formatted_address"`
	Geometry          Geometry            `json:"geometry"`
	PlaceID           string              `json:"place_id"`
	Types             []string            `json:"types"`
}
