package primer

import "sync"

// FactoryTableMutex is a mutex for the factory table
var FactoryTableMutex = sync.Mutex{}

var (
	FactoryPointer int64 = 0
	FactoryCursor  int64 = 0
	FactoryStep    int64 = 100
)
