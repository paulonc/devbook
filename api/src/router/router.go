package router

import (
	"api/src/controllers"
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func NewRouter(authContoller *controllers.AuthController, userController *controllers.UserController, publicationController *controllers.PublicationController) *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r, authContoller, userController, publicationController)
}
