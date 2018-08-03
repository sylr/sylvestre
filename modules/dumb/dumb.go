package main

import (
	"github.com/kataras/iris"
)

type SylvestreModule struct {
	name    string
	version string
}

func (module *SylvestreModule) Init() {
	module.name = "Dumb"
	module.version = "v0.0.1"
}

func (module SylvestreModule) Name() string {
	return module.name
}

func (module SylvestreModule) Version() string {
	return module.version
}

func (module SylvestreModule) RegisterEndpoint(httpServer *iris.Application) {
	httpServer.Get("/dumb", module.Dumb)
}

func (module SylvestreModule) Dumb(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "dumb",
	})
}

// exported module
var Module SylvestreModule
