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
		barf.Logger().Errorf(`[user.Register] [barf.Request(r).Body().Format(&data)] %s`, err.Error())
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Message: "We could not process your request at this time. Please try again later.",
		})
		return
	}

	user, err := userLogic.Register(&data)
	if err != nil {
		barf.Logger().Errorf(`[user.Register] [userLogic.Register(&data)] %s`, err.Error())
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	barf.Response(w).Status(http.StatusCreated).JSON(barf.Res{
		Status:  true,
		Data:    user,
		Message: "Registration successful.",
	})
}
