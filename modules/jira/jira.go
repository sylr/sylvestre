package main

import (
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
	"github.com/sylr/sylvestre/pkg"
)

type SylvestreModuleJira struct {
	name       string
	version    string
	coreConf   sylvestre.SylvestreCoreConfiguration
	moduleConf SylvestreModuleJiraConfiguration
}

type SylvestreModuleJiraConfiguration struct {
	Jira  JiraConfiguration  `yaml:"jira"`
	Slack SlackConfiguration `yaml:"slack"`
}

type JiraConfiguration struct {
	Token string `yaml:"token" conform:"redact"`
}

type SlackConfiguration struct {
	Token string `yaml:"token" conform:"redact"`
}

func (module *SylvestreModuleJira) Init() {
	module.name = "jira"
	module.version = "v0.0.1"
}

func (module *SylvestreModuleJira) SetConfiguration(coreConf *sylvestre.SylvestreCoreConfiguration, conf interface{}) {
	module.coreConf = *coreConf
	mapstructure.Decode(conf, &module.moduleConf)
}

func (module *SylvestreModuleJira) Name() string {
	return module.name
}

func (module *SylvestreModuleJira) Version() string {
	return module.version
}

func (module SylvestreModuleJira) RegisterEndpoint(httpServer *iris.Application) {
	httpServer.Get("/jira", module.Dispatch)
}

func (module SylvestreModuleJira) Dispatch(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": module.moduleConf,
	})
}

// exported module
var Module SylvestreModuleJira
