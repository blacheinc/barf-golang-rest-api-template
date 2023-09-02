package repository

import (
	"context"
	"fmt"

	"github.com/blacheinc/pixel/database"
	"github.com/blacheinc/pixel/primer"
)

type Factory struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

// GenerateAccountNumber generates a new account number for a new account
// by shifting the cursor by the given step. Account numbers are linearly
// random.
func GenerateAccountNumber() (string, error) {
	primer.FactoryTableMutex.Lock()
	defer primer.FactoryTableMutex.Unlock()
	if primer.FactoryCursor == 0 {
		cursor, err := ShiftCursorForKey(primer.FactoryStep, "account_number")
		if err != nil {
			return "", err
		}
		primer.FactoryCursor = cursor
		primer.FactoryPointer = cursor - primer.FactoryStep
	}
	if primer.FactoryPointer == primer.FactoryCursor {
		cursor, err := ShiftCursorForKey(primer.FactoryStep, "account_number")
		if err != nil {
			return "", err
		}
		primer.FactoryCursor = cursor
		primer.FactoryPointer = cursor - primer.FactoryStep
	}
	primer.FactoryPointer++
	return fmt.Sprintf("%010d", primer.FactoryPointer), nil
}

// ShiftCursorForAccountNumber shifts the cursor by the given step if the field exists else it creates it.
func ShiftCursorForKey(step int64, key string) (int64, error) {
	query := `INSERT INTO factory (key, value) VALUES ($1, $2) ON CONFLICT (key) DO UPDATE SET value = factory.value + $2 RETURNING value`
	var cursor int64
	err := database.PostgreSQLDB.NewRaw(query, key, step).Scan(context.Background(), &cursor)
	if err != nil {
		return 0, err
	}
	return cursor, nil
}
