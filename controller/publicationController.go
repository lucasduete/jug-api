package controller

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	"jug-api/model"
	"jug-api/dao/daoMongo"
	"jug-api/infraSecurity"
	dao2 "jug-api/dao"
	redis2 "github.com/go-redis/redis"
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
		redis := dao2.GetConnectionRedis()
		defer redis.Close()

		redis.Del("listPublications")
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
		redis := dao2.GetConnectionRedis()
		defer redis.Close()

		redis.Del("listPublications")
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
		redis := dao2.GetConnectionRedis()
		defer redis.Close()

		redis.Del("listPublications")
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

	publs := []model.Publication{}

	redis := dao2.GetConnectionRedis()
	defer redis.Close()

	cachePubls, err := redis.Get("listPublications").Result()
	if err != redis2.Nil {
		if err := json.Unmarshal([]byte(cachePubls), &publs); err != nil {
			return
		}

		respondWithJSON(response, 200, publs)
		return
	}

	dao := daoMongo.PublicationDaoMongo{}
	publs, err = dao.Listar()

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicações")
	} else if len(publs) == 0 {
		respondWithMessage(response, 204, "Não Há Publicações Salvas")
	} else {
		cachePubls, _ := json.Marshal(publs)
		redis.Set("listPublications", cachePubls, time.Hour*5).Result()

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
		return
	}

	dao := daoMongo.PublicationDaoMongo{}
	publs, err := dao.GetPublsByIndice(param)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
	} else if len(publs) == 0 {
		respondWithMessage(response, 204, "Não foram Encontradas Publicações")
	} else {
		respondWithJSON(response, 200, publs)
	}
}

func (app *App) GetRecomendation(response http.ResponseWriter, request *http.Request) {
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
		return
	}

	dao := daoMongo.PublicationDaoMongo{}
	publ, err := dao.GetPublById(idPublication)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
		return
	}

	publs, err := dao.GetRecomendation(publ.Titulo, publ.Tecnologia)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
		return
	}

	publsRecomendation := removeDuplicate(publs, idPublication)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
	} else if len(publsRecomendation) == 0 {
		respondWithMessage(response, 204, "Não foram Encontradas Publicações")
	} else {
		respondWithJSON(response, 200, publsRecomendation)
	}
}

func removeDuplicate(publs []model.Publication, index string) []model.Publication {

	publsClean := []model.Publication{}
	idPubl := bson.ObjectIdHex(index)
	for i := 0; i < len(publs); i++ {
		if strings.Compare(publs[i].ID.String(), idPubl.String()) != 0 {
			fmt.Println(publs[i].ID.String())
			publsClean = append(publsClean, publs[i])
		}
	}
	return publsClean
}
