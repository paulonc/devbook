package routes

import (
	"api/src/controllers"
	"net/http"
)

func PublicationRoutes(publicationController *controllers.PublicationController) []Route {
	return []Route{
		{
			URI:            "/publications",
			Method:         http.MethodPost,
			Function:       publicationController.CreatePublication,
			Authentication: true,
		},
		{
			URI:            "/publications",
			Method:         http.MethodGet,
			Function:       publicationController.GetPublications,
			Authentication: true,
		},
		{
			URI:            "/publications/{publicationId}",
			Method:         http.MethodGet,
			Function:       publicationController.GetPublication,
			Authentication: true,
		},
		{
			URI:            "/publications/{publicationId}",
			Method:         http.MethodPut,
			Function:       publicationController.UpdatePublication,
			Authentication: true,
		},
		{
			URI:            "/publications/{publicationId}",
			Method:         http.MethodDelete,
			Function:       publicationController.DeletePublication,
			Authentication: true,
		},
		{
			URI:            "/users/{userId}/publications",
			Method:         http.MethodGet,
			Function:       publicationController.SearchPublicationsByUser,
			Authentication: true,
		},
		{
			URI:            "/publications/{publicationId}/like",
			Method:         http.MethodPost,
			Function:       publicationController.LikePublication,
			Authentication: true,
		},
		{
			URI:            "/publications/{publicationId}/unlike",
			Method:         http.MethodPost,
			Function:       publicationController.UnlikePublication,
			Authentication: true,
		},
	}
}
