package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redmejia/models"
)

// update user business info email
func (a *App) BusinessHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPatch {
		var myBusinessInfo models.BusinessInformation
		err := json.NewDecoder(r.Body).Decode(&myBusinessInfo)
		if err != nil {
			a.ErrorLog.Fatal(err)
			return
		}

		a.InfoLog.Println("Update : ", myBusinessInfo)
		business := a.Db.GetBusinessInfoById(strconv.Itoa(myBusinessInfo.BusinessID)) // fix this

		if len(myBusinessInfo.Email) != 0 && myBusinessInfo.Email != business.Email {
			business.Email = myBusinessInfo.Email
		}

		var succes struct {
			Update       bool                       `json:"updated"`
			Error        bool                       `json:"error"`
			BusinessInfo models.BusinessInformation `json:"business_info"`
		}

		ok := a.Db.UpdateBusinessInfo(business)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			succes.Error = true
			succes.Update = false
			succes.BusinessInfo = models.BusinessInformation{}

			json.NewEncoder(w).Encode(&succes)
		}

		if ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			succes.Error = false
			succes.Update = true
			succes.BusinessInfo = *business

			json.NewEncoder(w).Encode(&succes)
		}

		a.InfoLog.Println(myBusinessInfo)

	} else if r.Method == http.MethodGet {

		// busID, err := strconv.Atoi(r.URL.Query().Get("bus_id"))
		busID := r.URL.Query().Get("bus_id")

		myBusiness := a.Db.GetBusinessInfoById(busID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(&myBusiness)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}

	}

}
