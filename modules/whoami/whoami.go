package main

import (
	"github.com/kataras/iris"
	"github.com/mitchellh/mapstructure"
	"github.com/nlopes/slack"
	"github.com/sylr/sylvestre/pkg"
)

type SylvestreModuleWhoAmI struct {
	name        string
	version     string
	coreConf    sylvestre.SylvestreCoreConfiguration
	moduleConf  SylvestreModuleWhoAmIConfiguration
	slackClient *slack.Client
}

type SylvestreModuleWhoAmIConfiguration struct {
	Slack SlackConfiguration `yaml:"slack"`
}

type SlackConfiguration struct {
	Token string `yaml:"token" conform:"redact"`
}

func (module *SylvestreModuleWhoAmI) Init() {
	module.name = "whoami"
	module.version = "v0.0.1"
}

func (module *SylvestreModuleWhoAmI) SetConfiguration(coreConf *sylvestre.SylvestreCoreConfiguration, conf interface{}) {
	module.coreConf = *coreConf
	mapstructure.Decode(conf, &module.moduleConf)
}

func (module *SylvestreModuleWhoAmI) Name() string {
	return module.name
}

func (module *SylvestreModuleWhoAmI) Version() string {
	return module.version
}

func (module SylvestreModuleWhoAmI) RegisterEndpoint(httpServer *iris.Application) {
	httpServer.Get("/whoami", module.Dispatch)
}

func (module SylvestreModuleWhoAmI) Dispatch(ctx iris.Context) {
	user_id := ctx.PostValue("user_id")
	user, _ := module.slackClient.GetUserProfile(user_id, true)

	ctx.JSON(iris.Map{
		"whoami": user,
	})
}

func (module *SylvestreModuleWhoAmI) getSlackClient() *slack.Client {
	if module.slackClient == nil {
		module.slackClient = slack.New(module.moduleConf.Slack.Token)
	}

	return module.slackClient
}

// exported module
var Module SylvestreModuleWhoAmI
