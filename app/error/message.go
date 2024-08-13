package error

type MessageError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *MessageError) Error() string {
	return e.Message
}

// NewMessageError creates a new MessageError
func NewMessageError(message string, code int) *MessageError {
	return &MessageError{
		Message: message,
		Code:    code,
	}
}
