package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/openapi-go/app/command"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"github.com/thanhhaudev/openapi-go/app/service"
	"github.com/thanhhaudev/openapi-go/app/util"

	"fmt"
	"net/http"
)

type messageHandler struct {
	messageService service.MessageService
	logger         *logrus.Logger
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
