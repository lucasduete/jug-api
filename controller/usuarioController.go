package controller

import (
	"jug-api/model"
	"jug-api/dao/daoPostgres"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func (app *App) SalvarUsuario(response http.ResponseWriter, request *http.Request) {
	user := model.User{}

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondWithMessage(response, 400, "Usuário Inválido")
		return
	}

	dao := daoPostgres.UserDaoPostgres{}
	err := dao.Salvar(user)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Salvar Usuário")
	} else {
		respondWithMessage(response, 200, "Usuário Cadastrado Com Sucesso")
	}
}

func (app *App) AtualizarUsuario(response http.ResponseWriter, request *http.Request) {

	user := model.User{}

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondWithMessage(response, 400, "Usuário Inválido")
	}

	dao := daoPostgres.UserDaoPostgres{}
	err := dao.Atualizar(user)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Atualizar Usuário")
	} else {
		respondWithMessage(response, 200, "Usuário Atualizado Com sucesso")
	}

}

func (app *App) RemoverUsuario(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]

	if email == "" {
		respondWithMessage(response, http.StatusBadRequest, "Email Inválido")
		return
	}

	dao := daoPostgres.UserDaoPostgres{}
	err := dao.Remover(email)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Remover Usuário")
	} else {
		respondWithMessage(response, 200, "Usuário Removido Com sucesso")
	}
}

func (app *App) ListarUsuarios(response http.ResponseWriter, request *http.Request) {

	dao := daoPostgres.UserDaoPostgres{}
	users, err := dao.Listar()

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Usuários")
	} else {
		respondWithJSON(response, 200, users)
	}
}

func (app *App) GetUserByEmail(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]

	if email == "" {
		respondWithMessage(response, 400, "Email Inválido")
		return
	}

	var user = model.User{}
	dao := daoPostgres.UserDaoPostgres{}

	err := dao.GetUserByEmail(user, email)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Usuário")
	} else {
		respondWithJSON(response, 200, user)
	}
}
