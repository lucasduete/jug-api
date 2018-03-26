package controller

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func (app *App) NotFound(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "ERROUUUUUU")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

