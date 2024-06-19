package routes

import (
	"api/src/controllers"
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Configure(r *mux.Router, authContoller *controllers.AuthController, userController *controllers.UserController, publicationController *controllers.PublicationController) *mux.Router {
	allRoutes := [][]Route{
		UserRoutes(userController),
		AuthRoutes(authContoller),
		PublicationRoutes(publicationController),
	}

	for _, routes := range allRoutes {
		for _, route := range routes {

			if route.Authentication {
				r.HandleFunc(route.URI,
					middlewares.Logger(middlewares.Authenticate(route.Function)),
				).Methods(route.Method)
			} else {
				r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
			}
		}
	}
	return r

}
