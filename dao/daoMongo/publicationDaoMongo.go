package daoMongo

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"jug-api/model"
	connection "jug-api/dao"
)

type PublicationDaoMongo struct{}

const collection_publ = "publication"

func (dao *PublicationDaoMongo) Salvar(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	publication.ID = bson.NewObjectId()
	err = conn.C(collection_publ).Insert(&publication)

	return err
}

func (dao *PublicationDaoMongo) Atualizar(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_publ).UpdateId(publication.ID, &publication)
	return err
}

func (dao *PublicationDaoMongo) Remover(publication model.Publication) error {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.C(collection_publ).Remove(&publication)
	return err
}

func (dao *PublicationDaoMongo) Listar() ([]model.Publication, error) {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

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

func (dao *PublicationDaoMongo) GetPublById(idPublication string) (*model.Publication, error) {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var publ model.Publication
	err = conn.C(collection_publ).FindId(bson.ObjectIdHex(idPublication)).One(&publ)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return &publ, nil
	}
}

func (dao *PublicationDaoMongo) GetPublsByTec(tecnology string) ([]model.Publication, error) {
	conn, err := connection.GetConnectionMongo()
	defer conn.Logout()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	publs := []model.Publication{}
	err = conn.C(collection_publ).Find(bson.M{"tecnologia": tecnology}).All(&publs)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return publs, nil
	}
}
