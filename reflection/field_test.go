package reflection

import (
	"testing"
	"time"
)

// go test -v -run TestReflectionUnit ./...
func TestReflectionUnit(t *testing.T) {

	t.Run("Should return the required struct fields", func(t *testing.T) {

		fields := ReturnStructFields(&struct {
			ID   int64
			Name string
			Time time.Time
			Skip string `rsf:"false"`
		}{})

		expected := 3

		if len(fields) != expected {
			t.Fatalf("unexpected number of fields: got %v want %v", len(fields), expected)
		}

	})

	t.Run("Should return the required struct fields with recursion", func(t *testing.T) {

		fields := ReturnStructFields(&struct {
			ID   int64
			Name string
			Time time.Time
			Skip string `rsf:"false"`
			Sub  struct {
				ID   int64
				Name string
			}
			SkipSub struct {
				ID   int64
				Name string
			} `rsfr:"false"` // skip recursion but return a single field as pointer to struct
		}{})

		expected := 6

		if len(fields) != expected {
			t.Fatalf("unexpected number of fields: got %v want %v", len(fields), expected)
		}

	})
}
