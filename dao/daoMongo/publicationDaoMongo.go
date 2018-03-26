package daoMongo

import (
	"jug-api/model"
	connection "jug-api/dao"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type publicationDaoMongo struct{}

const COLLECTION = "publication"

func (dao *publicationDaoMongo) Salvar(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(COLLECTION).Insert(&publication)
	return err
}

func (dao *publicationDaoMongo) Atualizar(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(COLLECTION).UpdateId(publication.ID, &publication)
	return err
}

func (dao *publicationDaoMongo) Remover(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(COLLECTION).Remove(&publication)
	return err
}

func (dao *publicationDaoMongo) Listar() ([]model.Publication, error) {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var publs = []model.Publication{}
	err = conn.C(COLLECTION).Find(bson.M{}).All(&publs)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return publs, nil
	}

}

func (dao *publicationDaoMongo) GetPublById(id int, pub model.Publication) (error) {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(COLLECTION).FindId(id).One(&pub)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}

}
