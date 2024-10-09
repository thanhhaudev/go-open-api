package handler

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/go-open-api/app/command"
	appErr "github.com/thanhhaudev/go-open-api/app/error"
	"github.com/thanhhaudev/go-open-api/app/repository"
	"github.com/thanhhaudev/go-open-api/app/service"
	"github.com/thanhhaudev/go-open-api/app/util"
	"strconv"

	"errors"
	"fmt"
	"net/http"
)

type messageHandler struct {
	messageService service.MessageService
	logger         *logrus.Logger
}

// GetMessage		godoc
// @Summary			Retrieve a message by its ID
// @Tags			message
// @Accept 			json
// @Produce			json
// @Param       	id path int true "Message ID"
// @Success     	200  {object} model.Message
// @Failure     	404  {object} error.MessageError
// @Failure     	400  {object} error.MessageError
// @Failure     	500  {object} error.MessageError
// @Router      	/api/v1/messages/{id} [get]
// @Security 		Bearer
func (m messageHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := m.messageService.FindMessageByID(uint(id))
	if err != nil {
		var notFound *appErr.MessageNotFoundError
		if errors.As(err, &notFound) {
			util.Response(w, err, http.StatusNotFound)

			return
		}

		util.Response(w, err, http.StatusBadRequest)

		return
	}

	util.Response(w, data, http.StatusOK)
}

// CreateMessage	godoc
// @Summary			Create a new message
// @Tags			message
// @Accept 			json
// @Produce			json
// @Param			request body command.MessageRequest true "Payload"
// @Success     	201  {object} model.Message
// @Failure     	404  {object} error.MessageError
// @Failure     	400  {object} error.MessageError
// @Failure     	500  {object} error.MessageError
// @Router      	/api/v1/messages [post]
// @Security 		Bearer
func (m messageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var com command.MessageRequest
	if err := util.Bind(r, &com); err != nil {
		m.logger.WithError(err).Error("Failed to parse request body")
		util.Response(w, err, http.StatusBadRequest)

		return
	}

	// validate request
	if err := com.Validate(); err != nil {
		fmt.Println(err)
		util.Response(w, err, http.StatusBadRequest)

		return
	}

	// create message
	data, err := m.messageService.CreateMessage(com)
	if err != nil {
		util.Response(w, err, http.StatusBadRequest)

		return
	}

	util.Response(w, data, http.StatusCreated)
}

// NewMessageHandler creates a new MessageHandler
func NewMessageHandler(
	userRepo repository.UserRepository,
	userMessageRepo repository.UserMessageRepository,
	messageRepo repository.MessageRepository,
	logger *logrus.Logger,
) MessageHandler {
	return &messageHandler{
		messageService: service.NewMessageService(userRepo, userMessageRepo, messageRepo, logger),
		logger:         logger,
	}
}
