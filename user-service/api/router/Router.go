package router

import (
	"net/http"

	"github.com/redmejia/api/handlers"
	"github.com/redmejia/middleware"
)

func Router(a *handlers.App) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1", a.AppHelthHandler)
	mux.HandleFunc("/v1/register", a.RegisterHandler)
	mux.HandleFunc("/v1/login", a.LoginHandler)
	mux.HandleFunc("/v1/deal", a.DealHandler)
	mux.HandleFunc("/v1/deals/all", a.DealsAllHandler)
	mux.HandleFunc("/v1/my/business", a.BusinessHandler)

	mux.HandleFunc("/v1/2024", a.HappyNewYearHandler)

	return middleware.Cors(mux)

}
