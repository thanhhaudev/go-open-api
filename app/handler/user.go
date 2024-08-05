package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	appErr "github.com/thanhhaudev/openapi-go/app/error"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"github.com/thanhhaudev/openapi-go/app/service"
	"github.com/thanhhaudev/openapi-go/app/util"
)

type userHandler struct {
	userService service.UserService
	logger      *logrus.Logger
}

// FindUser		godoc
// @Summary     Find user by ID
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       id path int true "User ID"
// @Success     200  {object} model.User
// @Failure     404  {object} error.UserError
// @Failure     400  {object} error.UserError
// @Failure     500  {object} error.UserError
// @Router      /api/users/{id} [get]
func (u userHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := u.userService.FindUserByID(uint(id))
	if err != nil {
		var notFound *appErr.UserError
		if errors.As(err, &notFound) {
			util.Response(w, err, http.StatusNotFound)

			return
		}

		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, data, http.StatusOK)
}

// GetUsers		godoc
// @Summary     Get all users
// @Tags        user
// @Accept      json
// @Produce     json
// @Success     200  {object} []model.User
// @Failure     400  {object} error.AuthError
// @Failure     500  {object} error.AuthError
// @Router      /api/users [get]
func (u userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	data, err := u.userService.GetUsers()
	if err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, data, http.StatusOK)
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(r repository.UserRepository, l *logrus.Logger) UserHandler {
	return &userHandler{
		userService: service.NewUserService(r, l),
		logger:      l,
	}
}
