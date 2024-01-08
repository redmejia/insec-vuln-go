package handlers

import (
	"fmt"
	"net/http"
)

func (a *App) HappyNewYearHandler(w http.ResponseWriter, r *http.Request) {
	happyNewYear := `
██   ██  █████  ██████  ██████  ██    ██     ███    ██ ███████ ██     ██     ██    ██ ███████  █████  ██████      ██████   ██████  ██████  ██   ██     
██   ██ ██   ██ ██   ██ ██   ██  ██  ██      ████   ██ ██      ██     ██      ██  ██  ██      ██   ██ ██   ██          ██ ██  ████      ██ ██   ██     
███████ ███████ ██████  ██████    ████       ██ ██  ██ █████   ██  █  ██       ████   █████   ███████ ██████       █████  ██ ██ ██  █████  ███████     
██   ██ ██   ██ ██      ██         ██        ██  ██ ██ ██      ██ ███ ██        ██    ██      ██   ██ ██   ██     ██      ████  ██ ██           ██     
██   ██ ██   ██ ██      ██         ██        ██   ████ ███████  ███ ███         ██    ███████ ██   ██ ██   ██     ███████  ██████  ███████      ██     
                                                                                                                                                    `

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, happyNewYear)
}
