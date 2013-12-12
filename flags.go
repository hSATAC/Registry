package main

import (
	"flag"
	"os"
)

var (
	ConfigFile = os.Getenv("ENVIRONMENT")
)

func init() {
	flag.StringVar(&ConfigFile, "config", "registry.yml", "Path to the Registry config file")
}

func Flags() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	return nil
}
