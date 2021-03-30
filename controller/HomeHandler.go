package controller

import (
	"encoding/json"
	"go-starter/model"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Status: "UP"})
}
