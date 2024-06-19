package controllers

import (
	"api/src/authentication"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PublicationController struct {
	repository repositories.PublicationRepository
}

func NewPublicationController(publicationRepository repositories.PublicationRepository) *PublicationController {
	return &PublicationController{repository: publicationRepository}
}

func (p *PublicationController) CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err := json.Unmarshal(body, &publication); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userID

	if err := publication.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	publication.ID, err = p.repository.CreatePublication(publication)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)

}
func (p *PublicationController) GetPublications(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	publications, err := p.repository.GetPublications(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

func (p *PublicationController) GetPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	publication, err := p.repository.GetPublication(publicationID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)

}

func (p *PublicationController) UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	savePublication, err := p.repository.GetPublication(publicationID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if savePublication.AuthorID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("it is not possible to update a post that is not yours"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err := json.Unmarshal(body, &publication); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := publication.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := p.repository.UpdatePublication(publicationID, publication); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func (p *PublicationController) DeletePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	savePublication, err := p.repository.GetPublication(publicationID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if savePublication.AuthorID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("it is not possible to delete a post that is not yours"))
		return
	}

	if err := p.repository.DeletePublication(publicationID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func (p *PublicationController) SearchPublicationsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	publications, err := p.repository.FindByUser(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, publications)

}

func (p *PublicationController) LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	savePublication, err := p.repository.GetPublication(publicationID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if savePublication.ID == 0 {
		responses.Err(w, http.StatusNotFound, errors.New("publication not found"))
		return
	}

	if err := p.repository.Like(publicationID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func (p *PublicationController) UnlikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	savePublication, err := p.repository.GetPublication(publicationID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if savePublication.ID == 0 {
		responses.Err(w, http.StatusNotFound, errors.New("publication not found"))
		return
	}

	if err := p.repository.Unlike(publicationID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
