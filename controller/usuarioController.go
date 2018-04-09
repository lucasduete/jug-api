package controller

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"jug-api/model"
	"jug-api/dao/daoPostgres"
	"jug-api/infraSecurity"
	"strings"
)

func (app *App) SalvarUsuario(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

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
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, email := infraSecurity.ValidateToken(token)

	if tokenValid == false || len(email) == 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	user := model.User{}

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondWithMessage(response, 400, "Usuário Inválido")
		return
	}

	if strings.Compare(user.Email, email) != 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
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
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, email := infraSecurity.ValidateToken(token)

	if tokenValid == false || len(email) == 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
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
	defer request.Body.Close()

	dao := daoPostgres.UserDaoPostgres{}
	users, err := dao.Listar()

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Usuários")
	} else if len(users) == 0 {
		respondWithMessage(response, 204, "Não há Usuários")
	} else {
		respondWithJSON(response, 200, users)
	}
}

func (app *App) GetUserByEmail(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	vars := mux.Vars(request)
	email := vars["email"]

	if len(email) == 0 {
		respondWithMessage(response, 400, "Email Inválido")
		return
	}

	dao := daoPostgres.UserDaoPostgres{}
	user, err := dao.GetUserByEmail(email)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Usuário")
	} else if len(user.Email) == 0 {
		respondWithMessage(response, 204, "Usuário não Existe")
	} else {
		respondWithJSON(response, 200, user)
	}
}

func (app *App) Login(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	senha := request.FormValue("senha")

	if email == "" || len(email) == 0 || senha == "" || len(senha) == 0 {
		respondWithMessage(response, 400, "Dados Inválidos")
	}

	dao := daoPostgres.UserDaoPostgres{}
	status, err := dao.Login(email, senha)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Realizar Login")
	} else if status == false {
		respondWithMessage(response, 401, "Usuário Não Autorizado")
	} else {
		user, _ := dao.GetUserByEmail(email)
		user.Senha, _ = infraSecurity.GenerateToken(email)
		respondWithJSON(response, 200, user)
	}
}
