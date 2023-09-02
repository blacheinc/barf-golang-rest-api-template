package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/blacheinc/pixel/database"
	"github.com/blacheinc/pixel/enum"
	"github.com/blacheinc/pixel/primer"
	"github.com/blacheinc/pixel/reflection"
	"github.com/blacheinc/pixel/types"
	"github.com/uptrace/bun/schema"
)

// Validate validates the user struct
func (u *User) Prepare() error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Age < primer.MinAge {
		return fmt.Errorf("age must be greater than %d", primer.MinAge-1)
	}
	u.Key = ""
	u.Email = strings.ToLower(u.Email)
	u.FirstName = strings.ToUpper(string(u.FirstName[0])) + strings.ToLower(u.FirstName[1:])
	u.LastName = strings.ToUpper(string(u.LastName[0])) + strings.ToLower(u.LastName[1:])
	return nil
}

/*
Date loads the created_at and updated_at fields of the activity if not already present, otherwise, it loads the updated_at field only.

If the "pessimistic" parameter is set to true, it loads both fields regardless
*/
func (u *User) Date(pessimistic ...bool) {
	if len(pessimistic) > 0 && !pessimistic[0] {
		if u.CreatedAt.IsZero() {
			u.CreatedAt = schema.NullTime{Time: time.Now()}
			u.UpdatedAt = schema.NullTime{Time: time.Now()}
			return
		}
		u.UpdatedAt = schema.NullTime{Time: time.Now()}
		return
	}
	u.CreatedAt = schema.NullTime{Time: time.Now()}
	u.UpdatedAt = schema.NullTime{Time: time.Now()}
}

/* Fields returns the struct fields as a slice of interface{} values */
func (u *User) Fields() []interface{} {
	return reflection.ReturnStructFields(u)
}

// Create inserts a new user into the database.
func (u *User) Create() error {
	if _, err := database.PostgreSQLDB.NewRaw(`INSERT INTO users (first_name, last_name, email, age, key, role, active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, u.FirstName, u.LastName, u.Email, u.Age, u.Key, u.Role, u.Active, u.CreatedAt, u.UpdatedAt).Exec(context.Background()); err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

/*
FByKeyVal finds and returns a user matching the key/value pair

# By default, only the id and name fields are loaded

The	"preloadandjoin" parameter can be used to request that all the fields of the struct be loaded

It returns an error if any
*/
func (u *User) FByKeyVal(key string, val interface{}, preloadandjoin ...bool) error {
	if len(preloadandjoin) > 1 && preloadandjoin[0] && preloadandjoin[1] {
		return database.PostgreSQLDB.NewRaw(`SELECT * FROM users LEFT JOIN configs ON users.id = configs.product_id WHERE users.`+key+` = ?`, val).Scan(context.Background(), u.Fields()...)
	}
	if len(preloadandjoin) > 0 && preloadandjoin[0] {
		return database.PostgreSQLDB.NewRaw(`SELECT * FROM users WHERE `+key+` = ?`, val).Scan(context.Background(), u)
	}
	return database.PostgreSQLDB.NewRaw(`SELECT id, firstName, lastName FROM users WHERE users.`+key+` = ?`, val).Scan(context.Background(), u)
}

/*
FByKeyVal finds and returns all users matching the key/value pair

# By default, only the id and name fields are loaded

The	"preloadandjoin" parameter can be used to request that all the fields of the struct be loaded

It returns an error if any
*/
func (u *Users) FByKeyVal(key string, val interface{}, limit, offset int, sort string, preloadandjoin ...bool) error {
	if len(preloadandjoin) > 1 && preloadandjoin[0] && preloadandjoin[1] {
		return database.PostgreSQLDB.NewRaw(`SELECT * FROM users LEFT JOIN configs ON users.id = configs.product_id WHERE users.`+key+` = ? ORDER BY users.updated_at `+sort+` LIMIT ? OFFSET ?`, val, limit, offset).Scan(context.Background(), u)
	}
	if len(preloadandjoin) > 0 && preloadandjoin[0] {
		return database.PostgreSQLDB.NewRaw(`SELECT * FROM users WHERE `+key+` = ? ORDER BY users.updated_at `+sort+` LIMIT ? OFFSET ?`, val, limit, offset).Scan(context.Background(), u)
	}
	return database.PostgreSQLDB.NewRaw(`SELECT id, name FROM users WHERE `+key+` = ? ORDER BY users.updated_at `+sort+` LIMIT ? OFFSET ?`, val, limit, offset).Scan(context.Background(), u)
}

/*
FByMap finds and returns a user matching the key/value pairs provided in the map

# By default, only the id and name fields are loaded

The	"preloadandjoin" parameter can be used to request that all the fields of the struct be loaded

It returns an error if any
*/
func (u *User) FByMap(m types.SQLMaps, preloadandjoin ...bool) error {
	query, args := database.MapsToWQuery(m)
	if len(preloadandjoin) > 1 && preloadandjoin[0] && preloadandjoin[1] {
		return database.PostgreSQLDB.NewRaw(`SELECT * FROM users LEFT JOIN configs ON users.id = configs.product_id WHERE `+query, args...).Scan(context.Background(), u.Fields()...)
	}
	if len(preloadandjoin) > 0 && preloadandjoin[0] {
		return database.PostgreSQLDB.NewRaw(`SELECT * FROM users WHERE `+query, args...).Scan(context.Background(), u)
	}
	return database.PostgreSQLDB.NewRaw(`SELECT id, firstName, lastName FROM users WHERE `+query, args...).Scan(context.Background(), u)
}

/*
FByMap finds and returns all users matching the key/value pairs provided in the map

# By default, only the id and name fields are loaded

The	"preloadandjoin" parameter can be used to request that all the fields of the struct be loaded

It returns an error if any
*/
func (u *Users) FByMap(m types.SQLMaps, limit, offset int, sort string, preloadandjoin ...bool) error {
	query, args := database.MapsToWQuery(m)
	if len(preloadandjoin) > 1 && preloadandjoin[0] && preloadandjoin[1] {
		rows, err := database.PostgreSQLDB.QueryContext(context.Background(), `SELECT * FROM users LEFT JOIN configs ON users.id = configs.product_id WHERE `+query+` ORDER BY users.updated_at `+sort+` LIMIT ? OFFSET ?`, append(args, limit, offset)...)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var user User
			if err := rows.Scan(user.Fields()...); err != nil {
				return err
			}
			*u = append(*u, user)
		}
	}
	if len(preloadandjoin) > 0 && preloadandjoin[0] {
		return database.PostgreSQLDB.NewRaw(`SELECT * FROM users WHERE `+query+` ORDER BY users.updated_at `+sort+` LIMIT ? OFFSET ?`, append(args, limit, offset)...).Scan(context.Background(), u)
	}
	return database.PostgreSQLDB.NewRaw(`SELECT id, firstName, lastName FROM users WHERE `+query+` ORDER BY users.updated_at `+sort+` LIMIT ? OFFSET ?`, append(args, limit, offset)...).Scan(context.Background(), u)
}

/*
UByMap updates a user matching the key/value pairs provided in the map

It returns an error if any
*/
func (u *User) UByMap(m types.SQLMaps) error {
	query, args := database.MapsToSQuery(m)
	if strings.Contains(query, string(enum.RETURNING)) {
		return database.PostgreSQLDB.NewRaw(`UPDATE users `+query, args...).Scan(context.Background(), u)
	}
	_, err := database.PostgreSQLDB.NewRaw(`UPDATE users `+query, args...).Exec(context.Background())
	return err
}

/*
CByMap finds and counts all users matching the key/value pairs provided in the map

It returns an error if any
*/
func (s *User) CByMap(m types.SQLMaps) (int, error) {
	var count int
	query, args := database.MapsToWQuery(m)
	err := database.PostgreSQLDB.NewRaw(`SELECT count(*) FROM savings WHERE `+query, args...).Scan(context.Background(), &count)
	return count, err
}

/*
CByMap finds and counts all users matching the key/value pairs provided in the map

It returns an error if any
*/
func (s *Users) CByMap(m types.SQLMaps) (int, error) {
	var count int
	query, args := database.MapsToWQuery(m)
	err := database.PostgreSQLDB.NewRaw(`SELECT count(*) FROM users WHERE `+query, args...).Scan(context.Background(), &count)
	return count, err
}

// Delete deletes a user from the database by email. This is only used for testing.
func (u *User) Delete() error {
	_, err := database.PostgreSQLDB.Exec(`DELETE FROM users WHERE email = $1`, u.Email)
	if err != nil {
		return err
	}
	return nil
}
