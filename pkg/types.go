package sylvestre

import (
	"github.com/kataras/iris"
)

// SylvestreOptions options
type SylvestreOptions struct {
	Configuration string `short:"c" long:"conf" description:"Configuration yaml file"`
	Verbose       []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Version       bool   `          long:"version" description:"Show version"`
}

// SylvestreCoreConfiguration
type SylvestreCoreConfiguration struct {
	ListeningAddresses []string `yaml:"listening-addresses"`
	ListeningPort      int      `yaml:"listening-port"`
	HTTPSInsecure      bool     `yaml:"https-insecure"`
	HTTPTimeout        int      `yaml:"http-timeout"`
}

// SylvestreCoreConfiguration
type SylvestreConfiguration struct {
	SylvestreCoreConfiguration `yaml:",inline"`
	Modules                    map[string]interface{} `yaml:"modules"`
}

// SylvestreModule interface
type SylvestreModule interface {
	Init()
	Name() string
	Version() string
	RegisterEndpoint(*iris.Application)
	SetConfiguration(*SylvestreCoreConfiguration, interface{})
}
