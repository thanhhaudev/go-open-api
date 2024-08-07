package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/openapi-go/app/command"
	appErr "github.com/thanhhaudev/openapi-go/app/error"
	"github.com/thanhhaudev/openapi-go/app/model"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"github.com/thanhhaudev/openapi-go/app/service"
	"github.com/thanhhaudev/openapi-go/app/util"
)

type userHandler struct {
	userService service.UserService
	logger      *logrus.Logger
}

// UpdateUser	godoc
// @Summary     Update a user
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body command.UserRequest true "Payload"
// @Success     201  {object} model.User
// @Failure     404  {object} error.UserError
// @Failure     400  {object} error.UserError
// @Failure     500  {object} error.UserError
// @Router      /api/v1/users/{id} [put]
// @Security 	Bearer
func (u userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// bind request
	com := command.UserRequest{}
	if err := util.Bind(r, &com); err != nil {
		u.logger.WithError(err).Error("Failed to parse request body")
		util.Response(w, err, http.StatusBadRequest)

		return
	}

	// validate request
	if err := com.Validate(); err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	user, err := u.userService.UpdateUser(uint(id), &com)
	if err != nil {
		var notFound *appErr.UserNotFoundError
		if errors.As(err, &notFound) {
			util.Response(w, err, http.StatusNotFound)

			return
		}

		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, user, http.StatusCreated)
}

// DeleteUser	godoc
// @Summary     Delete a user
// @Tags        user
// @Accept      json
// @Produce     json
// @Success     200  {object} map[string]bool
// @Failure     404  {object} error.UserError
// @Failure     400  {object} error.UserError
// @Failure     500  {object} error.UserError
// @Router      /api/v1/users/{id} [delete]
// @Security 	Bearer
func (u userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := u.userService.FindUserByID(uint(id))
	if err != nil {
		var notFound *appErr.UserNotFoundError
		if errors.As(err, &notFound) {
			util.Response(w, err, http.StatusNotFound)

			return
		}

		util.Response(w, err, http.StatusBadRequest)
		return
	}

	if err := u.userService.DeleteUser(user); err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, map[string]bool{"ok": true}, http.StatusOK)
}

// CreateUser	godoc
// @Summary     Create a new user
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       request body command.UserRequest true "Payload"
// @Success     201  {object} model.User
// @Failure     404  {object} error.UserError
// @Failure     400  {object} error.UserError
// @Failure     500  {object} error.UserError
// @Router      /api/v1/users [post]
// @Security 	Bearer
func (u userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	p := command.UserRequest{}
	if err := util.Bind(r, &p); err != nil {
		u.logger.WithError(err).Error("Failed to parse request body")
		util.Response(w, err, http.StatusBadRequest)

		return
	}

	// validate request
	if err := p.Validate(); err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	user, err := u.userService.CreateUser(&model.User{
		Email:       p.Email,
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
	})
	if err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, user, http.StatusCreated)
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
// @Router      /api/v1/users/{id} [get]
// @Security 	Bearer
func (u userHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := u.userService.FindUserByID(uint(id))
	if err != nil {
		var notFound *appErr.UserNotFoundError
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
// @Summary     Retrieve all users
// @Tags        user
// @Accept      json
// @Produce     json
// @Success     200  {object} []model.User
// @Failure     400  {object} error.AuthError
// @Failure     500  {object} error.AuthError
// @Router      /api/v1/users [get]
// @Security 	Bearer
func (u userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	data, err := u.userService.GetUsers()
	if err != nil {
		util.Response(w, err, http.StatusBadRequest)
		return
	}

	util.Response(w, data, http.StatusOK)
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(
	userRepo repository.UserRepository,
	userMessageRepo repository.UserMessageRepository,
	logger *logrus.Logger,
) UserHandler {
	return &userHandler{
		userService: service.NewUserService(userRepo, userMessageRepo, logger),
		logger:      logger,
	}
}
