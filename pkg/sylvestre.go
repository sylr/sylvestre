package sylvestre

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"

	"gopkg.in/yaml.v2"

	"github.com/kataras/iris"
	log "github.com/sirupsen/logrus"
)

type Sylvestre struct {
	httpServer    *iris.Application
	Configuration SylvestreConfiguration
}

func (app *Sylvestre) Init() {
	app.httpServer = iris.Default()
}

func (app *Sylvestre) LoadConfiguration() {
	app.Configuration = app.parseConfigFile("conf/sylvestre.yml")
}

func (app Sylvestre) parseConfigFile(file string) SylvestreConfiguration {
	conf := SylvestreConfiguration{}

	content, _ := ioutil.ReadFile(file)

	yaml.Unmarshal(content, &conf)

	return conf
}

func (app *Sylvestre) GetCoreConfiguration() *SylvestreCoreConfiguration {
	return &SylvestreCoreConfiguration{
		ListeningAddresses: app.Configuration.ListeningAddresses,
		ListeningPort:      app.Configuration.ListeningPort,
		HTTPSInsecure:      app.Configuration.HTTPSInsecure,
		HTTPTimeout:        app.Configuration.HTTPTimeout,
	}
}

func (app *Sylvestre) Run() {
	app.httpServer.Run(iris.Addr(":8080"))
}

func (app *Sylvestre) LoadModules() {
	modfiles, _ := filepath.Glob("modules/*/*.so")

	for _, modfile := range modfiles {
		log.Infof("Modules: loading '%s'", modfile)

		plug, err := plugin.Open(modfile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		log.Debugf("Modules: module '%s' loaded", modfile)

		symModule, err := plug.Lookup("Module")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var module SylvestreModule
		module, ok := symModule.(SylvestreModule)

		if !ok {
			log.Errorf("Modules: module found in '%s' does not implement SylvestreModule", modfile)
			os.Exit(1)
		}

		module.Init()
		module.SetConfiguration(app.GetCoreConfiguration(), app.Configuration.Modules[module.Name()])
		module.RegisterEndpoint(app.httpServer)

		log.Infof("Modules: module %s:%s initialized", module.Name(), module.Version())
	}
}
