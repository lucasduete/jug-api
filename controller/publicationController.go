package controller

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"jug-api/model"
	"jug-api/dao/daoMongo"
	"time"
	"jug-api/infraSecurity"
	"strings"
)

func (app *App) SalvarPublication(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, email := infraSecurity.ValidateToken(token)

	if tokenValid == false || len(email) == 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	publ := model.Publication{}

	if err := json.NewDecoder(request.Body).Decode(&publ); err != nil {
		respondWithMessage(response, 400, "Publicação Inválida")
		return
	}

	publ.Data = time.Now()
	publ.EmailUser = email

	dao := daoMongo.PublicationDaoMongo{}
	err := dao.Salvar(publ)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Salvar Publicação")
	} else {
		respondWithMessage(response, 200, "Publicação Salva com Sucesso")
	}
}

func (app *App) AtualizarPublication(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	publ := model.Publication{}

	if err := json.NewDecoder(request.Body).Decode(&publ); err != nil {
		respondWithMessage(response, 400, "Publicação Inválida")
	}

	dao := daoMongo.PublicationDaoMongo{}
	err := dao.Atualizar(publ)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Atualizar Publicação")
	} else {
		respondWithMessage(response, 200, "Publicação Atualizada com Sucesso")
	}
}

func (app *App) RemoverPublication(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, email := infraSecurity.ValidateToken(token)

	if tokenValid == false || len(email) == 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	publ := model.Publication{}

	if err := json.NewDecoder(request.Body).Decode(&publ); err != nil {
		respondWithMessage(response, 400, "Publicação Inválida")
	}

	if strings.Compare(publ.EmailUser, email) != 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	dao := daoMongo.PublicationDaoMongo{}
	err := dao.Remover(publ)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Remover Publicação")
	} else {
		respondWithMessage(response, 200, "Publicação Removida com Sucesso")
	}
}

func (app *App) ListarPublications(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	dao := daoMongo.PublicationDaoMongo{}
	publs, err := dao.Listar()

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicações")
	} else if len(publs) == 0 {
		respondWithMessage(response, 204, "Não Há Publicações Salvas")
	} else {
		respondWithJSON(response, 200, publs)
	}
}

func (app *App) GetPublById(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	vars := mux.Vars(request)
	idPublication := vars["idPublication"]

	if len(idPublication) == 0 {
		respondWithMessage(response, 400, "ID Inválido")
	}

	dao := daoMongo.PublicationDaoMongo{}
	publ, err := dao.GetPublById(idPublication)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
	} else if publ.Conteudo == "" {
		respondWithMessage(response, 204, "Publicação Não Foi Econtrada")
	} else {
		respondWithJSON(response, 200, publ)
	}

}

func (app *App) GetPublsByTec(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	vars := mux.Vars(request)
	tecnology := vars["tecnologia"]

	if tecnology == "" || len(tecnology) == 0 {
		respondWithMessage(response, 400, "Tecnologia Inválida")
	}

	dao := daoMongo.PublicationDaoMongo{}
	publs, err := dao.GetPublsByTec(tecnology)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
	} else if len(publs) == 0 {
		respondWithMessage(response, 204, "Publicação Não Foi Econtrada")
	} else {
		respondWithJSON(response, 200, publs)
	}
}

func (app *App) GetPublsByIndice(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	param := request.FormValue("param")

	if len(param) == 0 {
		respondWithMessage(response, 400, "Parametro inválido.")
	}

	dao := daoMongo.PublicationDaoMongo{}
	publs, err := dao.GetPublsByIndice(param)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
	} else if len(publs) == 0 {
		respondWithMessage(response, 204, "Publicação Não Foi Econtrada")
	} else {
		respondWithJSON(response, 200, publs)
	}

}