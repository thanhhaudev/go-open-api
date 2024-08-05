package error

import "net/http"

type UserError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (u UserError) Error() string {
	return u.Message
}

// NewUserError creates a new UserError with the given message and code
func NewUserError(message string, code int) *UserError {
	return &UserError{
		Message: message,
		Code:    code,
	}
}

// NewUserNotFoundError creates a new UserError with message "User not found" and code 404
func NewUserNotFoundError() *UserError {
	return NewUserError("User not found", http.StatusNotFound)
}
