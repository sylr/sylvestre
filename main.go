// vim: ts=4:

package main

import (
	"fmt"
	"os"

	"github.com/sylr/sylvestre/pkg"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var (
	opts    = sylvestre.SylvestreOptions{}
	parser  = flags.NewParser(&opts, flags.Default)
	version = "v0.0.1"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// main
func main() {
	// looing for --version in args
	for _, val := range os.Args {
		if val == "--version" {
			fmt.Printf("sylvestre version %s\n", version)
			os.Exit(0)
		} else if val == "--" {
			break
		}
	}

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	// Update logging level
	switch {
	case len(opts.Verbose) >= 1:
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.Debugf("Options: %+v", opts)
	log.Infof("Version: %s", version)

	syl := &sylvestre.Sylvestre{}
	syl.Init()
	syl.LoadConfiguration()
	syl.LoadModules()
	syl.Run()
}
