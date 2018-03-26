package daoPostgres

import (
	"jug-api/model"
	connection "jug-api/dao"
	"log"
)

type UserDaoPostgres struct{}

func (dao *UserDaoPostgres) Salvar(user model.User) error {
	conn, err := connection.GetConnection()

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = conn.Exec("INSERT INTO User(Nome, Username, Email, Senha) "+
		"VALUES ($1,$2,$3,$4", user.Nome, user.Username, user.Email, user.Senha)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}

}

func (dao *UserDaoPostgres) Atualizar(user model.User) error {
	conn, err := connection.GetConnection()

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = conn.Exec("UPDATE User SET Nome = $1, Username = $2, Senha = $3 "+
		"WHERE Email LIKE $4", user.Nome, user.Username, user.Senha, user.Email)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}

}

func (dao *UserDaoPostgres) Remover(email string) error {
	conn, err := connection.GetConnection()

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = conn.Exec("DELETE FROM User WHERE Email like $1", email)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}
}

func (dao *UserDaoPostgres) Listar() ([]model.User, error) {
	conn, err := connection.GetConnection()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := conn.Query("SELECT Nome, Username, Email, Senha  FROM User")

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

func (dao *UserDaoPostgres) GetUserByEmail(user model.User, email string) (error) {
	conn, err := connection.GetConnection()

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.QueryRow("SELECT Nome, Username, Email, Senha  FROM User "+
		"WHERE Email LIKE $1", email).Scan(&user.Nome, &user.Username, &user.Email, &user.Senha,
		&user.Email)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}

}
