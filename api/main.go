package main

import (
	"api/src/config"
	"api/src/server"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.Load()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
