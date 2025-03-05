package handlers

import (
	"encoding/json"
	"geo_service/internal/models"
	"net/http"
)

func GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var geocodeReq models.GeocodeRequest

	if err := json.NewDecoder(r.Body).Decode(&geocodeReq); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	client := createDaDataClient()

	addresses, err := client.GeocodeAddress(geocodeReq.Lat, geocodeReq.Lng)
	if err != nil {
		handleAPIError(w, err)
		return
	}

	geocodeResp := models.GeocodeResponse{Addresses: convertToPointerSlice(addresses)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(geocodeResp)
}
