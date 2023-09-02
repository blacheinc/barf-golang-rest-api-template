package user

import (
	userController "github.com/blacheinc/pixel/controller/v1/user"
	"github.com/opensaucerer/barf"
)

func RegisterUserRoutes(frame *barf.SubRoute) {

	frame = frame.RetroFrame("/accounts")

	frame.Post("/register", userController.Register)
}
