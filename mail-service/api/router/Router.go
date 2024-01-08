package router

import (
	"net/http"

	"github.com/redmejia/mail/api/handler"
	"github.com/redmejia/middleware"
)

func Router(a *handler.App) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/mail", a.HelthCheckHandler)
	mux.HandleFunc("/v1/mail/send", a.MailHandler)

	return middleware.Cors(mux)
}
