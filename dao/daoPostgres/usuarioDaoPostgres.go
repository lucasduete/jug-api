package daoPostgres

import (
	"log"
	"strings"

	"jug-api/model"
	connection "jug-api/dao"
)

type UserDaoPostgres struct{}

func (dao *UserDaoPostgres) Salvar(user model.User) error {
	conn, err := connection.GetConnectionPostgres()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = conn.Exec("INSERT INTO Usuario(Nome, Username, Email, Senha) "+
		"VALUES ($1,$2,$3,$4)", user.Nome, user.Username, user.Email, user.Senha)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}

}

func (dao *UserDaoPostgres) Atualizar(user model.User) error {
	conn, err := connection.GetConnectionPostgres()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = conn.Exec("UPDATE Usuario SET Nome = $1, Username = $2, Senha = $3 "+
		"WHERE Email LIKE $4", user.Nome, user.Username, user.Senha, user.Email)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}

}

func (dao *UserDaoPostgres) Remover(email string) error {
	conn, err := connection.GetConnectionPostgres()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = conn.Exec("DELETE FROM Usuario WHERE Email like $1", email)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}
}

func (dao *UserDaoPostgres) Listar() ([]model.User, error) {
	conn, err := connection.GetConnectionPostgres()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := conn.Query("SELECT Nome, Username, Email, Senha  FROM Usuario")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	users := []model.User{}

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.Nome, &user.Username, &user.Email, &user.Senha); err != nil {
			log.Fatal(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (dao *UserDaoPostgres) GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	conn, err := connection.GetConnectionPostgres()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = conn.QueryRow("SELECT Nome, Username, Email, Senha FROM Usuario "+
		"WHERE Email LIKE $1", email).Scan(&user.Nome, &user.Username, &user.Email, &user.Senha)

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return &user, nil
	}
}

func (dao *UserDaoPostgres) Login(email, senha string) (bool, error) {
	conn, err := connection.GetConnectionPostgres()

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	var password string
	err = conn.QueryRow("SELECT Senha FROM Usuario WHERE Email LIKE $1", email).Scan(&password)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	if strings.Compare(senha, password) == 0 {
		return true, nil
	} else {
		return false, nil
	}

}
