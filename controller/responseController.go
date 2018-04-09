package controller

import (
	"strconv"
	"net/http"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	"jug-api/model"
	"jug-api/dao/daoMongo"
	"time"
	"jug-api/infraSecurity"
	"strings"
)

func (app *App) SalvarResposta(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, email := infraSecurity.ValidateToken(token)

	if tokenValid == false || len(email) == 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	resp := model.Response{}

	if err := json.NewDecoder(request.Body).Decode(&resp); err != nil {
		respondWithMessage(response, 400, "Resposta Inválida")
		return
	}

	resp.Data = time.Now()
	resp.EmailUser = email

	dao := daoMongo.ResponseDaoMongo{}
	err := dao.Salvar(resp)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Cadastrar Resposta")
	} else {
		respondWithMessage(response, 200, "Resposta Cadastrada com Sucesso")
	}
}

func (app *App) AtualizarResposta(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, email := infraSecurity.ValidateToken(token)

	if tokenValid == false || len(email) == 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	resp := model.Response{}

	if err := json.NewDecoder(request.Body).Decode(&resp); err != nil {
		respondWithMessage(response, 400, "Resposta Inválida")
	}

	if strings.Compare(resp.EmailUser, email) != 0 {
		respondWithMessage(response, http.StatusUnauthorized, "")
		return
	}

	dao := daoMongo.ResponseDaoMongo{}
	err := dao.Atualizar(resp)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Atualizar Resposta")
	} else {
		respondWithMessage(response, 200, "Resposta Atualizada com Sucesso")
	}
}

func (app *App) RemoverResposta(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, email := infraSecurity.ValidateToken(token)

	if tokenValid == false || len(email) == 0 {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	resp := model.Response{}

	if err := json.NewDecoder(request.Body).Decode(&resp); err != nil {
		respondWithMessage(response, 400, "Resposta Inválida")
		return
	}

	if err := json.NewDecoder(request.Body).Decode(&resp); err != nil {
		respondWithMessage(response, 400, "Resposta Inválida")
	}


	dao := daoMongo.ResponseDaoMongo{}
	err := dao.Remover(resp)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Remover Resposta")
	} else {
		respondWithMessage(response, 200, "Resposta Removida com Sucesso")
	}
}

func (app *App) ListarRespostas(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	dao := daoMongo.ResponseDaoMongo{}
	resps, err := dao.Listar()

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Respostas")
	} else if len(resps) == 0 {
		respondWithMessage(response, 204, "Não Há Respostas Cadastradas")
	} else {
		respondWithJSON(response, 200, resps)
	}
}

func (app *App) GetRespById(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil || id <= 0 {
		respondWithMessage(response, 400, "ID Inválido")
	}

	resp := model.Response{}
	dao := daoMongo.ResponseDaoMongo{}

	err = dao.GetRespById(id, resp)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Respostas")
	} else if resp.Conteudo == "" {
		respondWithMessage(response, 204, "Não Há Respostas Cadastradas")
	} else {
		respondWithJSON(response, 200, resp)
	}
}

func (app *App) GetRespByPubl(response http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("Authorization")
	tokenValid, _ := infraSecurity.ValidateToken(token)

	if tokenValid == false {
		respondWithMessage(response, http.StatusUnauthorized, "Token Inválido")
		return
	}

	vars := mux.Vars(request)
	temp := vars["idPublication"]

	if temp == "" || len(temp) == 0 {
		respondWithMessage(response, 400, "Id da Publucação é Inválida")
	}

	idPublication := bson.ObjectIdHex(temp)
	dao := daoMongo.ResponseDaoMongo{}

	resps, err := dao.GetRespsByPubl(idPublication)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Respostas")
	} else if len(resps) == 0 {
		respondWithMessage(response, 204, "Não Há Respostas Cadastradas")
	} else {
		respondWithJSON(response, 200, resps)
	}

}
