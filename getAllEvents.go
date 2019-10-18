package main

import (
	"encoding/json"
	"net/http"
)

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}
