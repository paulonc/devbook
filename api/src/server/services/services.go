package services

import (
	"api/src/controllers"
	"api/src/repositories"
	"database/sql"
)

type Services struct {
	AuthController        *controllers.AuthController
	UserController        *controllers.UserController
	PublicationController *controllers.PublicationController
}

func Initialize(db *sql.DB) (*Services, error) {
	userRepository := repositories.NewUserRepository(db)
	publicationRepository := repositories.NewPublicationRepository(db)

	userController := controllers.NewUserController(userRepository)
	authContoller := controllers.NewAuthController(userRepository)
	publicationController := controllers.NewPublicationController(publicationRepository)

	return &Services{
		AuthController:        authContoller,
		UserController:        userController,
		PublicationController: publicationController,
	}, nil
}
