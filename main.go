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

	//Users EndPoints
	router.HandleFunc("/api/usuarios", app.SalvarUsuario).Methods("POST")
	router.HandleFunc("/api/usuarios", app.AtualizarUsuario).Methods("PATCH")
	router.HandleFunc("/api/usuarios", app.RemoverUsuario).Methods("DELETE")
	router.HandleFunc("/api/usuarios", app.ListarUsuarios).Methods("GET")
	router.HandleFunc("/api/usuarios/usuario", app.GetUserByEmail).Methods("GET")

	//Defaults EndPoints
	router.HandleFunc("/", app.NotFound).Methods("GET")

	fmt.Println("Servidor Rodando na Porta " + port)
	http.ListenAndServe(":"+port, router)
}
