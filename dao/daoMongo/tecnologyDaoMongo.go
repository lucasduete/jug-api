package daoMongo

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"jug-api/model"
	connection "jug-api/dao"
)

type TecnologyDaoMongo struct{}

const collection_tec = "tecnology"

func (dao *TecnologyDaoMongo) Salvar(tecnology model.Tecnology) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	tecnology.ID = bson.NewObjectId()
	err = conn.C(collection_tec).Insert(&tecnology)

	return err
}

func (dao *TecnologyDaoMongo) Atualizar(id int, tecnology model.Tecnology) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_tec).UpdateId(id, &tecnology)
	return err
}

func (dao *TecnologyDaoMongo) Remover(tecnology model.Tecnology) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_tec).Remove(&tecnology)
	return err
}

func (dao *TecnologyDaoMongo) Listar() ([]model.Tecnology, error) {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var tecs = []model.Tecnology{}
	err = conn.C(collection_tec).Find(bson.M{}).All(&tecs)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return tecs, err
	}
}
