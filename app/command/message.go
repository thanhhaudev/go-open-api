package command

import "fmt"

type MessageRequest struct {
	SenderId    uint   `json:"sender_id" example:"1"`
	ReceiverIds []uint `json:"receiver_ids" example:"2"`
	Subject     string `json:"subject" example:"Hello!"`
	Content     string `json:"content" example:"Hello, how are you?"`
}

// Validate validates the message request
func (m *MessageRequest) Validate() error {
	if m.SenderId == 0 {
		return fmt.Errorf("sender_id is required")
	}

	if len(m.ReceiverIds) == 0 {
		return fmt.Errorf("receiver_ids is required")
	}

	if m.Subject == "" {
		return fmt.Errorf("subject is required")
	}

	if m.Content == "" {
		return fmt.Errorf("content is required")
	}

	return nil
}
