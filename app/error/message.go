package error

type MessageError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *MessageError) Error() string {
	return e.Message
}

type MessageNotFoundError MessageError

func (e *MessageNotFoundError) Error() string {
	return e.Message
}

// NewMessageNotFoundError creates a new MessageNotFoundError
func NewMessageNotFoundError() *MessageNotFoundError {
	return &MessageNotFoundError{
		Message: "Message not found",
		Code:    404,
	}
}

// NewMessageError creates a new MessageError
func NewMessageError(message string, code int) *MessageError {
	return &MessageError{
		Message: message,
		Code:    code,
	}
}
