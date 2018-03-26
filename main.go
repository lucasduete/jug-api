package main

import (
	"os"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	if port == "" {
		port = "8080"
	}

	router.HandleFunc("/", notFound).Methods("GET")

	fmt.Println("Servidor Rodando na Porta " + port)
	http.ListenAndServe(":"+port, router)
}

func notFound(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "ERROUUUUUU")
}
