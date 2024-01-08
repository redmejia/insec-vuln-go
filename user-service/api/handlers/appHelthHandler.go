package handlers

import (
	"encoding/json"
	"net/http"
)

func (a *App) AppHelthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}
}
