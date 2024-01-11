package main

import (
	"github.com/joho/godotenv"
	"userservice/pkg/repository"
	"userservice/pkg/server"
)

import (
	"fmt"
	"github.com/jinzhu/configor"
)

type Config struct {
	Server   server.Config
	Postgres repository.Config
}

func Parse(cfg any) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("parse config %v", err)
	}
	configorLoader := configor.New(&configor.Config{
		Silent:               true,
		ErrorOnUnmatchedKeys: true,
		Environment:          "",
		ENVPrefix:            "-",
		Debug:                false,
		Verbose:              false,
		AutoReload:           false,
		AutoReloadInterval:   0,
		AutoReloadCallback:   nil,
	})

	if err := configorLoader.Load(cfg); err != nil {
		return fmt.Errorf("load config %v", err)
	}

	return nil
}