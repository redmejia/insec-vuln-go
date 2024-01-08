package handler

import "net/http"

func (a *App) HelthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mail Service v1"))
}
