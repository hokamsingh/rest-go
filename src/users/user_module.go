package user

import (
	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

func NewUserModule() *LessGo.Module {
	userService := NewUserService()
	userController := NewUserController(userService, "/users")
	return LessGo.NewModule("UserModule",
		[]interface{}{userController}, // Controllers
		[]interface{}{userService},    // Services
	)
}
