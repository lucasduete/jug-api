package controller

import (
	"net/http"
	"jug-api/dao/daoMongo"
	"encoding/json"
	"jug-api/model"
	"github.com/gorilla/mux"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"encoding/hex"
)

func (app *App) SalvarResposta(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	resp := model.Response{}

	if err := json.NewDecoder(request.Body).Decode(&resp); err != nil {
		respondWithMessage(response, 400, "Resposta Inválida")
		return
	}

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

	resp := model.Response{}

	if err := json.NewDecoder(request.Body).Decode(&resp); err != nil {
		respondWithMessage(response, 400, "Resposta Inválida")
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

	resp := model.Response{}

	if err := json.NewDecoder(request.Body).Decode(&resp); err != nil {
		respondWithMessage(response, 400, "Resposta Inválida")
		return
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
