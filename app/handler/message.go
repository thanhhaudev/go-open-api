package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/openapi-go/app/service"
)

type messageHandler struct {
	service service.MessageService
	logger  *logrus.Logger
}

func NewMessageHandler(logger *logrus.Logger) MessageHandler {
	return &messageHandler{
		logger: logger,
	}
}
