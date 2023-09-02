package types

import (
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
