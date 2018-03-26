package controller

import (
	"jug-api/model"
	"jug-api/dao/daoPostgres"
)

func (app *App) SalvarUsuario (user model.User) error {
	dao := daoPostgres.UserDaoPostgres{}
	return dao.Salvar(user)
}

func (app *App) AtualizarUsuario (user model.User) error {
	dao := daoPostgres.UserDaoPostgres{}
	return dao.Atualizar(user)
}

func (app *App) RemoverUsuario (email string) error {
	dao := daoPostgres.UserDaoPostgres{}
	return dao.Remover(email)
}

func (app *App) ListarUsuarios() ([]model.User, error) {
	dao := daoPostgres.UserDaoPostgres{}

	return dao.Listar()
}

func (app *App) GetUserByEmail (email string) (model.User, error) {
	var user = model.User{}
	dao := daoPostgres.UserDaoPostgres{}

	err := dao.GetUserByEmail(user, email)

	return user, err
}
