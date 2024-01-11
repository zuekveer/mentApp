package main

import (
	"log"
	"userservice/pkg/repository"
	"userservice/pkg/runner"
	"userservice/pkg/server"
)

func main() {
	var cfg Config
	err := Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(&cfg.Server)
	if err := srv.Start(); err != nil {
		return
	}

	repo := repository.New(&cfg.Postgres)
	if err := repo.Start(); err != nil {
		return
	}

	rnr := runner.New(&r.Run)
	if err := rnr.Start(); err != nil {
		return
	}




}