package main

import (
	"github.com/kataras/iris"
)

type SylvestreModule struct {
	name    string
	version string
}

func (module *SylvestreModule) Init() {
	module.name = "Pwet"
	module.version = "v0.0.1"
}

func (module SylvestreModule) Name() string {
	return module.name
}

func (module SylvestreModule) Version() string {
	return module.version
}

func (module SylvestreModule) RegisterEndpoint(httpServer *iris.Application) {
	httpServer.Get("/pwet", module.Pwet())
}

func (module SylvestreModule) Pwet() func(iris.Context) {
	return func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pwet",
		})
	}
}

// exported module
var Module SylvestreModule
