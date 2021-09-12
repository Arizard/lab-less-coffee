package main

import (
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/core"
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/infra"
)

func main() {
	server := infra.NewGinServer()
	core := core.Core{
		Server:      server,
		UserService: entity.NewUserService(infra.NewJSONUserRepository("./infra/data/user.json")),
	}

	core.Start()

}
