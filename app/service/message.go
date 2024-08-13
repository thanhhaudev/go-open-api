package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/openapi-go/app/command"
	appErr "github.com/thanhhaudev/openapi-go/app/error"
	"github.com/thanhhaudev/openapi-go/app/model"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"gorm.io/gorm"

	"errors"
	"fmt"
	"net/http"
)

type (
	MessageService interface {
		CreateMessage(com command.MessageRequest) (*model.Message, error)
		FindMessageByID(id uint) (*model.Message, error)
	}

	messageService struct {
		userRepository        repository.UserRepository
		userMessageRepository repository.UserMessageRepository
		messageRepository     repository.MessageRepository
		redisClient           *redis.Client
		logger                *logrus.Logger
	}
)

// FindMessageByID finds a message by its ID
func (m messageService) FindMessageByID(id uint) (*model.Message, error) {
	message, err := m.messageRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.NewMessageNotFoundError()
		}

		return nil, err
	}

	return message, nil
}

// CreateMessage creates a new message
func (m messageService) CreateMessage(com command.MessageRequest) (*model.Message, error) {
	// validate sender
	sender, err := m.userRepository.FindByID(com.SenderId)
	if err != nil {
		logrus.WithError(err).Error("Failed to find sender")

		return nil, appErr.NewMessageError("Failed to find sender", http.StatusBadRequest)
	}

	// validate receivers
	receivers, err := m.userRepository.FindByIDs(com.ReceiverIds)
	if err != nil {
		return nil, err
	}

	// Create a map of found receiver IDs
	foundReceiverIDs := make(map[uint]bool)
	for _, receiver := range receivers {
		foundReceiverIDs[receiver.ID] = true
	}

	// prevent sending message to self
	if foundReceiverIDs[com.SenderId] {
		return nil, appErr.NewMessageError(fmt.Sprintf("Cannot send message to self: %d", com.SenderId), http.StatusBadRequest)
	}

	// Identify invalid receiver IDs
	var invalidIDs []uint
	for _, id := range com.ReceiverIds {
		if !foundReceiverIDs[id] {
			invalidIDs = append(invalidIDs, id)
		}
	}

	if len(invalidIDs) > 0 {
		return nil, appErr.NewMessageError(fmt.Sprintf("%s not exist: %v", func() string {
			if len(invalidIDs) == 1 {
				return "user"
			}

			return "users"
		}(), invalidIDs), http.StatusBadRequest)
	}

	// Create a new message
	message := &model.Message{
		Sender:    sender,
		Receivers: receivers,
		Subject:   com.Subject,
		Content:   com.Content,
	}

	// Save the message
	if err := m.messageRepository.Create(message); err != nil {
		return nil, err
	}

	return message, nil
}

// NewMessageService creates a new MessageService
func NewMessageService(
	userRepo repository.UserRepository,
	userMessageRepo repository.UserMessageRepository,
	messageRepo repository.MessageRepository,
	logger *logrus.Logger,
) MessageService {
	return &messageService{
		userRepository:        userRepo,
		userMessageRepository: userMessageRepo,
		messageRepository:     messageRepo,
		logger:                logger,
	}
}
