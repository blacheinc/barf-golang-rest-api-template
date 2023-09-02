package primer

type AccountType int
type Type int
type Status int

const (
	_ AccountType = iota

	Savings

	Current
)

const (
	_ Type = iota

	Deposit

	Withdrawal

	Lock

	Unlock
)

const (
	Pending Status = iota

	Completed
)
