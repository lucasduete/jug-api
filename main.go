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
	url_base := "/api/"
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	if port == "" {
		port = "8080"
	}

	//Users EndPoints
	router.HandleFunc(url_base+"usuarios/", app.SalvarUsuario).Methods("POST")
	router.HandleFunc(url_base+"usuarios/", app.AtualizarUsuario).Methods("PATCH")
	router.HandleFunc(url_base+"usuarios/{email}", app.RemoverUsuario).Methods("DELETE")
	router.HandleFunc(url_base+"usuarios/", app.ListarUsuarios).Methods("GET")
	router.HandleFunc(url_base+"usuarios/usuario/{email}", app.GetUserByEmail).Methods("GET")
	router.HandleFunc(url_base+"usuarios/login/", app.Login).Methods("POST")

	//Tecnology EndPoints
	router.HandleFunc(url_base+"tecnologias/", app.SalvarTecnologia).Methods("POST")
	router.HandleFunc(url_base+"tecnologias/", app.AtualizarTecnologia).Methods("PATCH")
	router.HandleFunc(url_base+"tecnologias/", app.RemoverTecnologia).Methods("DELETE")
	router.HandleFunc(url_base+"tecnologias/", app.ListarTecnologias).Methods("GET")

	//Response EndPoints
	router.HandleFunc(url_base+"responses/", app.SalvarResposta).Methods("POST")
	router.HandleFunc(url_base+"responses/", app.AtualizarResposta).Methods("PATCH")
	router.HandleFunc(url_base+"responses/", app.RemoverResposta).Methods("DELETE")
	router.HandleFunc(url_base+"responses/", app.ListarTecnologias).Methods("GET")
	router.HandleFunc(url_base+"responses/response/", app.GetRespById).Methods("POST")
	router.HandleFunc(url_base+"responses/publication/", app.GetRespByPubl).Methods("POST")

	//Publication EndPoints
	router.HandleFunc(url_base+"publications/", app.SalvarPublication).Methods("POST")
	router.HandleFunc(url_base+"publications/", app.AtualizarPublication).Methods("PATCH")
	router.HandleFunc(url_base+"publications/", app.RemoverPublication).Methods("DELETE")
	router.HandleFunc(url_base+"publications/", app.ListarPublications).Methods("GET")
	router.HandleFunc(url_base+"publications/publication/", app.GetPublById).Methods("POST")
	router.HandleFunc(url_base+"publications/tecnology/", app.GetPublsByTec).Methods("POST")

	//Defaults EndPoints
	router.HandleFunc("/*", app.NotFound)

	fmt.Println("Servidor Rodando na Porta " + port)
	http.ListenAndServe(":"+port, router)
}
