package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/openapi-go/app/repository"
)

type (
	MessageService interface {
		// TODO
	}

	messageService struct {
		messageRepository repository.MessageRepository
		redisClient       *redis.Client
		logger            *logrus.Logger
	}
)

// NewMessageService creates a new MessageService
func NewMessageService(
	messageRepository repository.MessageRepository,
	redisClient *redis.Client,
	logger *logrus.Logger,
) MessageService {
	return &messageService{
		messageRepository: messageRepository,
		redisClient:       redisClient,
		logger:            logger,
	}
}
