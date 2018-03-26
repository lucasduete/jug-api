package main

import (
	"os"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"jug-api/model"
)

var app = model.App{}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()


	if port == "" {
		port = "8080"
	}

	router.HandleFunc("/", app.notFound).Methods("GET")

	fmt.Println("Servidor Rodando na Porta " + port)
	http.ListenAndServe(":"+port, router)
}

func (app *) notFound(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "ERROUUUUUU")
}
