package daoMongo

import (
	"jug-api/model"
	connection "jug-api/dao"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type PublicationDaoMongo struct{}

const collection_publ = "publication"

func (dao *PublicationDaoMongo) Salvar(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_publ).Insert(&publication)
	return err
}

func (dao *PublicationDaoMongo) Atualizar(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_publ).UpdateId(publication.ID, &publication)
	return err
}

func (dao *PublicationDaoMongo) Remover(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_publ).Remove(&publication)
	return err
}

func (dao *PublicationDaoMongo) Listar() ([]model.Publication, error) {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var publs = []model.Publication{}
	err = conn.C(collection_publ).Find(bson.M{}).All(&publs)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return publs, nil
	}

}

func (dao *PublicationDaoMongo) GetPublById(id int, pub model.Publication) (error) {
	conn, err := connection.GetConnectionMongo()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_publ).FindId(id).One(&pub)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}

}