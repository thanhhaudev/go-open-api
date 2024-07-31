package handler

import (
	"github.com/thanhhaudev/openapi-go/app/config"
	"github.com/thanhhaudev/openapi-go/app/datastore/mysql"
	"github.com/thanhhaudev/openapi-go/app/service"
)

type userHandler struct {
	userService service.UserService
}

func (u userHandler) GetUser() string {
	//TODO implement me
	panic("implement me")
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(db *config.Database) UserHandler {
	r := mysql.NewUserRepository(db.Conn)

	return &userHandler{
		userService: service.NewUserService(r),
	}
}
