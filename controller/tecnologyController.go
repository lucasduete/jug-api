package controller

import (
	"net/http"
	"jug-api/model"
	"encoding/json"
	"jug-api/dao/daoMongo"
	"log"
	"github.com/gorilla/mux"
	"strconv"
)

func (app *App) SalvarTecnologia(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	tec := model.Tecnology{}

	if err := json.NewDecoder(request.Body).Decode(&tec); err != nil {
		respondWithMessage(response, 400, "Usuário Inválido")
		return
	}

	dao := daoMongo.TecnologyDaoMongo{}
	err := dao.Salvar(tec)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Cadastrar Tecnologia")
	} else {
		respondWithMessage(response, 200, "Tecnologia Cadastrada com Sucesso")
	}
}

func (app *App) AtualizarTecnologia(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	tec := model.Tecnology{}

	if err := json.NewDecoder(request.Body).Decode(&tec); err != nil {
		log.Fatal(err)
		return
	}

	dao := daoMongo.TecnologyDaoMongo{}
	err := dao.Salvar(tec)

	if err != nil {
		log.Fatal(err)
		respondWithMessage(response, 500, "Erro ao Atualizar Tecnologia")
	} else {
		respondWithMessage(response, 200, "Tecnologia Atualizada com Sucesso")
	}

}

func (app *App) RemoverTecnologia(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	tec := model.Tecnology{}

	if err := json.NewDecoder(request.Body).Decode(&tec); err != nil {
		respondWithMessage(response, 400, "Tecnologia Inválida")
		return
	}

	dao := daoMongo.TecnologyDaoMongo{}
	err := dao.Remover(tec)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Remover Tecnologia")
	} else {
		respondWithMessage(response, 200, "Tecnologia Removida com Sucesso")
	}
}

func (app *App) ListarTecnologias(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	dao := daoMongo.TecnologyDaoMongo{}
	tecs, err := dao.Listar()

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Tecnologias")
	} else if len(tecs) == 0 {
		respondWithMessage(response, 204, "Não Há Tecnologias Cadastradas")
	} else {
		respondWithJSON(response, 200, tecs)
	}
}

func (app *App) GetTecById(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil || id <= 0 {
		respondWithMessage(response, 400, "Id Inválido")
	}

	tec := model.Tecnology{}
	dao := daoMongo.TecnologyDaoMongo{}

	if err := dao.GetTecById(id, tec); err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Tecnologia")
	} else if tec.Nome == "" {
		respondWithMessage(response, 204, "Tecnologia Não Encontrada")
	} else {
		respondWithJSON(response, 200, tec)
	}

}