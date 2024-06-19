package server

import (
	"api/src/config"
	"api/src/database"
	"api/src/migrations"
	"api/src/router"
	"api/src/server/services"
	"fmt"
	"log"
	"net/http"
)

func Start() error {
	db, err := database.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer db.Close()

	err = migrations.RunMigrations()
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	s, err := services.Initialize(db)
	if err != nil {
		return fmt.Errorf("failed to initialize services: %w", err)
	}

	r := router.NewRouter(s.AuthController, s.UserController, s.PublicationController)

	log.Printf("Listening on port %d\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}
