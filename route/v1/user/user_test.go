package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	userController "github.com/blacheinc/pixel/controller/v1/user"
	userRepository "github.com/blacheinc/pixel/repository/v1/user"
	"github.com/blacheinc/pixel/test"
	"github.com/opensaucerer/barf"
)

// go test -v -run TestUserRouteIntegration ./...
func TestUserRouteIntegration(t *testing.T) {

	test.Setup()

	data := userRepository.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
		Age:       30,
	}
	// convert struct to bytes
	datab, _ := json.Marshal(data)

	t.Run("Should create a new user", func(t *testing.T) {

		req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(datab))
		if err != nil {
			t.Fatal(err)
		}

		writer := httptest.NewRecorder()
		handler := http.HandlerFunc(userController.Register)

		handler.ServeHTTP(writer, req)

		if status := writer.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		// convert response body to struct
		var res barf.Res
		if err := json.NewDecoder(writer.Body).Decode(&res); err != nil {
			t.Fatal(err)
		}

		if res.Status != true {
			t.Fatalf("handler returned unexpected status: got %v want %v",
				res.Status, true)
		}

		if res.Data.(map[string]interface{})["first_name"] != data.FirstName {
			t.Fatalf("handler returned unexpected first name: got %v want %v",
				res.Data.(map[string]interface{})["first_name"], data.FirstName)
		}
	})

	t.Run("Should fail to create a new user", func(t *testing.T) {

		req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(datab))
		if err != nil {
			t.Fatal(err)
		}

		writer := httptest.NewRecorder()
		handler := http.HandlerFunc(userController.Register)

		handler.ServeHTTP(writer, req)

		if status := writer.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// convert response body to struct
		var res barf.Res
		if err := json.NewDecoder(writer.Body).Decode(&res); err != nil {
			t.Fatal(err)
		}

		if res.Status != false {
			t.Fatalf("handler returned unexpected status: got %v want %v",
				res.Status, false)
		}

	})

	// clean up
	data.Delete()
}
