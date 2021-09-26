package core

import "github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"

type Core struct {
	UserService        *entity.UserService
	Server             Server
	UserSessionManager entity.UserSessionManager
}

func (core *Core) Start() {
	core.Server.Start(core)
}

func (core *Core) RegisterUser(login string, password string) (user *entity.User, err error) {
	reg := entity.UserRegistration{
		Login:    login,
		Password: password,
	}

	user, err = core.UserService.Register(reg)

	return user, err
}

func (core *Core) GetUserByLogin(login string) (*entity.User, error) {
	repo := core.UserService.UserRepository()

	return repo.GetByLogin(login)
}

func (core *Core) GetUserByUID(uid entity.UserUID) (*entity.User, error) {
	repo := core.UserService.UserRepository()

	return repo.Get(uid)
}
