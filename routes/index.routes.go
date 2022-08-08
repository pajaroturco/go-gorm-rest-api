package routes

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("Primer api rest")
}
