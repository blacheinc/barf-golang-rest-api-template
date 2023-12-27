package primitive

import "strings"

type String string

// ToSentence converts a string to sentence case (first letter capitalized)
func (s String) ToSentence() string {
	if len(s) < 1 {
		return ""
	}
	return strings.ToUpper(string(s[0])) + string(s[1:])
}

// String returns the string value
func (s String) String() string {
	return string(s)
}

// ToLower converts a string to lowercase
func (s String) ToLower() string {
	return strings.ToLower(s.String())
}

// ToUpper converts a string to uppercase
func (s String) ToUpper() string {
	return strings.ToUpper(s.String())
}
