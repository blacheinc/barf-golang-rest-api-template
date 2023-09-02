package user

import (
	"net/http"

	userLogic "github.com/blacheinc/pixel/logic/v1/user"
	userRepository "github.com/blacheinc/pixel/repository/v1/user"
	"github.com/opensaucerer/barf"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var data userRepository.User
	if err := barf.Request(r).Body().Format(&data); err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	user, err := userLogic.Register(&data)
	if err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	barf.Response(w).Status(http.StatusCreated).JSON(barf.Res{
		Status:  true,
		Data:    user,
		Message: "user created successfully",
	})
}
