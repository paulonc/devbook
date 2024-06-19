package routes

import (
	"api/src/controllers"
	"net/http"
)

func AuthRoutes(authController *controllers.AuthController) []Route {
	return []Route{
		{
			URI:            "/login",
			Method:         http.MethodPost,
			Function:       authController.Login,
			Authentication: false,
		},
	}
}
