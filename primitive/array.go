package primitive

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

// I have a plan to make a package from this called nanites. Basically, nanites will expose multi primitives implemented in the Go way. https://github.com/opensaucerer/nanites
type StringArray []string

// ExistsIn reports whether any of the elements of sa is contained in t.
func (sa StringArray) ExistsIn(t string) bool {
	for _, v := range sa {
		if strings.Contains(t, v) {
			return true
		}
	}
	return false
}

type Array []interface{}

// Len returns the length of the array
func (a Array) Len() int {
	return len(a)
}

// Includes returns true if the array includes the provided value
func (a Array) Includes(val interface{}) bool {
	for _, v := range a {
		if v == val {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the provided value in the array, or -1 if not found
func (a Array) IndexOf(val interface{}) int {
	for i, v := range a {
		if v == val {
			return i
		}
	}
	return -1
}

// ExistsIn reports whether t contains any of the elements of a that is a valid string.
// None string elements in a are ignored during the check.
func (a Array) ExistsIn(t string) bool {
	for _, v := range a {
		if _, ok := v.(string); !ok {
			continue
		}
		if strings.Contains(t, v.(string)) {
			return true
		}
	}
	return false
}

// Scan implements the Scanner interface.
func (sa *StringArray) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, sa)
	case string:
		return json.Unmarshal([]byte(v), sa)
	case nil:
		return nil
	}
	return nil
}

// Value implements the driver Valuer interface.
func (sa StringArray) Value() (driver.Value, error) {
	b, err := json.Marshal(sa)
	return string(b), err
}
