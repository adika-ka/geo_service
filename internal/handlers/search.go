package handlers

import (
	"encoding/json"
	"geo_service/internal/models"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var searchReq models.SearchRequest

	if err := json.NewDecoder(r.Body).Decode(&searchReq); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	client := createDaDataClient()

	addresses, err := client.SearchAddress(searchReq.Query)
	if err != nil {
		handleAPIError(w, err)
		return
	}

	searchResp := models.SearchResponse{Addresses: convertToPointerSlice(addresses)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResp)
}
