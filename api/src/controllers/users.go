package controllers

import (
	"api/src/authentication"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	repository repositories.UserRepository
}

func NewUserController(repository repositories.UserRepository) *UserController {
	return &UserController{repository: repository}
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare("registration"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	user.ID, err = c.repository.CreateUser(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func (c UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := c.repository.GetUsers()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func (c UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.repository.GetUser(ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)

}

func (u UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if ID != tokenUserID {
		responses.Err(w, http.StatusForbidden, errors.New("you cannot update a user other than yourself"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare("update"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := u.repository.UpdateUser(ID, user); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func (u UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if ID != tokenUserID {
		responses.Err(w, http.StatusForbidden, errors.New("you cannot delete a user other than yourself"))
		return
	}

	if err := u.repository.DeleteUser(ID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

}

func (u *UserController) FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if ID == followerID {
		responses.Err(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}

	if err := u.repository.FollowUser(ID, followerID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func (u *UserController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if ID == followerID {
		responses.Err(w, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
		return
	}

	if err := u.repository.UnfollowUser(ID, followerID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func (u *UserController) GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	followers, err := u.repository.GetFollowers(ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

func (u *UserController) GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	followers, err := u.repository.GetFollowing(ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

func (u *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if tokenID != ID {
		responses.Err(w, http.StatusForbidden, errors.New("you cannot update the password other than yourself"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password
	if err := json.Unmarshal(body, &password); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	storedPassword, err := u.repository.GetPassword(ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.ValidatePassword(storedPassword, password.Current); err != nil {
		responses.Err(w, http.StatusUnauthorized, errors.New("password wrong"))
		return
	}

	hashedPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := u.repository.UpdatePassword(ID, string(hashedPassword)); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
