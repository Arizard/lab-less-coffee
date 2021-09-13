package main

import (
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/core"
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/infra"
)

func main() {
	userRepo := infra.NewMySQLUserRepository()
	server := infra.NewGinServer()
	core := core.Core{
		Server:      server,
		UserService: entity.NewUserService(userRepo),
	}

	core.Start()

}
