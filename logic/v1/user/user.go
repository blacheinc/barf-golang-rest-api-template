package user

import (
	"errors"

	"github.com/blacheinc/pixel/primer"
	userRepository "github.com/blacheinc/pixel/repository/v1/user"
)

// Register registers a new user
func Register(user *userRepository.User) (*userRepository.User, error) {

	// this is not good enough, we can improve this with a validation middleware. Ideally, I would create one from scratch as most validators have used in golang are just not sufficient for me. I have a plan to work on one here https://github.com/opensaucerer/vibranium
	if err := user.Prepare(); err != nil {
		return nil, err
	}

	// ensure email is unique
	if err := user.FByKeyVal("email", user.Email); err != nil {
		return nil, errors.New("we are having issues verifying this email address. Please try again later")
	}

	if user.Key != "" {
		return nil, errors.New("a user with this email address already exists")
	}

	user.Role = primer.User
	user.Active = false

	// create user
	if err := user.Create(); err != nil {
		return nil, errors.New("we are having issues creating your account. Please try again later")
	}

	return user, nil
}
