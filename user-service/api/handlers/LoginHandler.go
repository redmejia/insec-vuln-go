package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redmejia/models"
)

func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// check on db for email
		var userLogin models.Login
		err := json.NewDecoder(r.Body).Decode(&userLogin)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}

		var succesfulLogin SuccessResponse

		user := a.Db.NewLogIn(userLogin)
		if user.Email == "" && user.Password == "" {

			a.ErrorLog.Println("user not found")

			succesfulLogin.Message = "DENIED"
			succesfulLogin.Success = false
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			err := json.NewEncoder(w).Encode(&succesfulLogin)
			if err != nil {
				a.ErrorLog.Fatal(err)
			}
		}

		if len(user.Email) > 0 && len(user.Password) > 0 {

			myBusiness := a.Db.GetBusinessInfoById(strconv.Itoa(user.BusinessID))
			a.InfoLog.Println("My buss ", myBusiness)

			succesfulLogin.Message = "GRANTED"
			succesfulLogin.Success = true
			succesfulLogin.Business = *myBusiness

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			err = json.NewEncoder(w).Encode(&succesfulLogin)
			if err != nil {
				a.ErrorLog.Fatal(err)
			}
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
}
