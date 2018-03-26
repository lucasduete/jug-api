package daoMongo

import (
	"jug-api/model"
	connection "jug-api/dao"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type TecnologyDaoMongo struct{}

const collection_tec = "tecnology"

func (dao *TecnologyDaoMongo) Salvar(tecnology model.Tecnology) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_tec).Insert(&tecnology)
	return err
}

func (dao *TecnologyDaoMongo) Atualizar(id int, tecnology model.Tecnology) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_tec).UpdateId(id, &tecnology)
	return err
}

func (dao *TecnologyDaoMongo) Remover(tecnology model.Tecnology) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_tec).Remove(&tecnology)
	return err
}

func (dao *TecnologyDaoMongo) Listar() ([]model.Tecnology, error) {
	conn, err := connection.GetConnectionMongo()

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

func (dao *TecnologyDaoMongo) GetTecById(id int, tecnology model.Tecnology) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_tec).FindId(id).One(&tecnology)
	return err
}