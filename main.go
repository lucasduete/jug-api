package main

import (
	"os"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"jug-api/controller"
)

var app = controller.App{}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()


	if port == "" {
		port = "8080"
	}

	router.HandleFunc("/", app.NotFound).Methods("GET")

	fmt.Println("Servidor Rodando na Porta " + port)
	http.ListenAndServe(":"+port, router)
}
