package types

import (
	"testing"
)

// go test -v -run TestTypesUnit ./...
func TestTypesUnit(t *testing.T) {

	t.Run("Should confirm that string exists in array", func(t *testing.T) {

		fields := StringArray{"one", "two", "three"}

		expected := true

		if fields.ExistsIn("two") != expected {
			t.Fatalf("unexpected result: got %v want %v", fields.ExistsIn("two"), expected)
		}

	})
}
