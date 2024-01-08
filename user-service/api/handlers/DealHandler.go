package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/models"
)

func (a *App) DealHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var deal models.Deal
		err := json.NewDecoder(r.Body).Decode(&deal)
		if err != nil {
			a.ErrorLog.Fatal(err)
			return
		}

		ok := a.Db.CreateNewDeal(deal)
		var newDeal struct {
			Create bool `json:"created"`
			Error  bool `json:"error"`
		}
		if !ok {
			newDeal.Create = false
			newDeal.Error = true
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)

			err := json.NewEncoder(w).Encode(&newDeal)
			if err != nil {
				a.ErrorLog.Fatal(err)
				return
			}
		}

		newDeal.Create = true
		newDeal.Error = false
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(&newDeal)
		if err != nil {
			a.ErrorLog.Fatal(err)
			return
		}
	} else if r.Method == http.MethodPatch {

		var deal models.Deal

		json.NewDecoder(r.Body).Decode(&deal)

		dealInfo := a.Db.GetDealByDealID(deal.DealID)
		a.InfoLog.Println(dealInfo)

		if len(deal.ProductName) != 0 && deal.ProductName != dealInfo.ProductName {
			dealInfo.ProductName = deal.ProductName
		}

		if len(deal.ProductDescription) != 0 && deal.ProductDescription != dealInfo.ProductDescription {
			dealInfo.ProductDescription = deal.ProductDescription
		}

		if deal.Price > 0 && deal.Price != dealInfo.Price {
			dealInfo.Price = deal.Price
		}

		dealInfo.DealID = deal.DealID
		ok := a.Db.UpdateDeal(dealInfo)
		var updated struct {
			Updated bool        `json:"updated"`
			Error   bool        `json:"error"`
			Deal    models.Deal `json:"deal"`
		}
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			updated.Updated = false
			updated.Error = true
			updated.Deal = models.Deal{}

			json.NewEncoder(w).Encode(&updated)

		}

		if ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			updated.Updated = true
			updated.Error = false
			updated.Deal = *dealInfo

			json.NewEncoder(w).Encode(&updated)
		}

	} else if r.Method == http.MethodGet {
		busName := r.URL.Query().Get("bus_name")

		name := a.Db.GetDealByBusinessName(busName)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(&name)
		if err != nil {
			a.ErrorLog.Fatal(err)
			return
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

}

// all deals
func (a *App) DealsAllHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		list := a.Db.GetAllDeals()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(&list)
		if err != nil {
			a.ErrorLog.Fatal(err)
			return
		}
	}

}
