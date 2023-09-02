package user

import (
	"testing"

	userRepository "github.com/blacheinc/pixel/repository/v1/user"
	"github.com/blacheinc/pixel/test"
)

// go test -v -run TestUserLogicIntegration ./...
func TestUserLogicIntegration(t *testing.T) {

	test.Setup()

	data := userRepository.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
		Age:       30,
	}

	t.Run("Should create a new user", func(t *testing.T) {

		u, err := Register(&data)
		if err != nil {
			t.Fatal(err)
		}

		if u.Key == "" {
			t.Fatalf("key should not be empty: got %v", u.Key)
		}

		if u.FirstName != data.FirstName {
			t.Fatalf("logic returned unexpected first name: got %v want %v",
				u.FirstName, data.FirstName)
		}

	})

	// clean up
	data.Delete()
}
