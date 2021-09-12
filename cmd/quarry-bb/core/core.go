package core

import "github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"

type Core struct {
	UserService *entity.UserService
	Server      Server
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
