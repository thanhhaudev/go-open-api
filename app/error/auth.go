package error

import "net/http"

type AuthError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e AuthError) Error() string {
	return e.Message
}

// NewAuthError creates a new AuthError instance with the given message and status code
func NewAuthError(message string, code int) *AuthError {
	return &AuthError{
		Message: message,
		Code:    code,
	}
}

// NewUnauthorizedError creates a new AuthError instance with status code 401
func NewUnauthorizedError() *AuthError {
	return NewAuthError("Unauthorized", http.StatusUnauthorized)
}

func NewForbiddenError() *AuthError {
	return NewAuthError("Forbidden", http.StatusForbidden)
}
