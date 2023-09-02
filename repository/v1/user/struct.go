package user

import (
	"github.com/blacheinc/pixel/primer"
	"github.com/uptrace/bun"
)

type User struct {
	Id        int64        `json:"-"`
	Key       string       `json:"key"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Age       int          `json:"age"`
	Role      primer.Role  `json:"role"`
	Active    bool         `json:"active"`
	CreatedAt bun.NullTime `json:"created_at"`
	UpdatedAt bun.NullTime `json:"updated_at"`
}

type Users []User
