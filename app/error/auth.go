package error

type AuthError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e AuthError) Error() string {
	return e.Message
}
