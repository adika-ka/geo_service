package models

// Маршрут: /api/address/search метод POST
type SearchRequest struct {
	Query string `json:"query"`
}

// Маршрут: /api/address/geocode метод POST
type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
