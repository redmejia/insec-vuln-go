package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/redmejia/models"
)

type SuccessResponse struct {
	Success  bool                       `json:"success"`
	Message  string                     `json:"message"`
	Business models.BusinessInformation `json:"my_business"`
}

func (a *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var userRegister models.Register
		err := json.NewDecoder(r.Body).Decode(&userRegister)
		if err != nil {
			log.Fatal(err)
		}

		ok, busID := a.Db.RegisterNewUser(userRegister)

		if !ok && busID == 0 {
			successResp := SuccessResponse{
				Success:  false,
				Message:  "DENIED",
				Business: models.BusinessInformation{},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			err = json.NewEncoder(w).Encode(&successResp)
			if err != nil {
				log.Fatal(err)
			}
		}

		if ok && busID > 1 {

			myBusiness := a.Db.GetBusinessInfoById(strconv.Itoa(busID))
			successResp := SuccessResponse{
				Success:  true,
				Message:  "GRANTED",
				Business: *myBusiness,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			err = json.NewEncoder(w).Encode(&successResp)

			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

}
