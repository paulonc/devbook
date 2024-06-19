package routes

import (
	"api/src/controllers"
	"net/http"
)

func UserRoutes(userController *controllers.UserController) []Route {
	return []Route{
		{
			URI:            "/users",
			Method:         http.MethodPost,
			Function:       userController.CreateUser,
			Authentication: false,
		},
		{
			URI:            "/users",
			Method:         http.MethodGet,
			Function:       userController.GetUsers,
			Authentication: true,
		},
		{
			URI:            "/users/{id}",
			Method:         http.MethodGet,
			Function:       userController.GetUser,
			Authentication: true,
		},
		{
			URI:            "/users/{id}",
			Method:         http.MethodPut,
			Function:       userController.UpdateUser,
			Authentication: true,
		},
		{
			URI:            "/users/{id}",
			Method:         http.MethodDelete,
			Function:       userController.DeleteUser,
			Authentication: true,
		},
		{
			URI:            "/users/{id}/follow",
			Method:         http.MethodPost,
			Function:       userController.FollowUser,
			Authentication: true,
		},
		{
			URI:            "/users/{id}/unfollow",
			Method:         http.MethodPost,
			Function:       userController.UnfollowUser,
			Authentication: true,
		},
		{
			URI:            "/users/{id}/followers",
			Method:         http.MethodGet,
			Function:       userController.GetFollowers,
			Authentication: true,
		},
		{
			URI:            "/users/{id}/following",
			Method:         http.MethodGet,
			Function:       userController.GetFollowing,
			Authentication: true,
		},
		{
			URI:            "/users/{id}/password",
			Method:         http.MethodPost,
			Function:       userController.UpdatePassword,
			Authentication: true,
		},
	}
}
