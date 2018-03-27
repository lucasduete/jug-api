package daoMongo

import (
	"jug-api/model"
	connection "jug-api/dao"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type ResponseDaoMongo struct {}

const collection_resp = "response"

func (dao *ResponseDaoMongo) Salvar(resp model.Response) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_resp).Insert(&resp)

	return err
}

func (dao *ResponseDaoMongo) Atualizar(resp model.Response) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_resp).UpdateId(resp.ID, &resp)

	return nil
}

func (dao *ResponseDaoMongo) Remover(resp model.Response) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_resp).Remove(&resp)
	return err

}

func (dao *ResponseDaoMongo) Listar() ([]model.Response, error)  {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var resps = []model.Response{}
	err = conn.C(collection_resp).Find(bson.M{}).All(&resps)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return resps, nil
	}
}

func (dao *ResponseDaoMongo) GetRespById(id int, response model.Response) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_resp).FindId(id).One(&response)
	return err
}

func (dao *ResponseDaoMongo) GetRespsByPubl(idPublication bson.ObjectId) ([]model.Response, error) {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var resps = []model.Response{}
	err = conn.C(collection_resp).Find(bson.M{"idPublication": idPublication}).All(&resps)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return resps, nil
	}
}