package enum

type SQLOperator string

const (
	// Equal is the SQL operator for equality
	Equal SQLOperator = "="

	// NotEqual is the SQL operator for inequality
	NotEqual SQLOperator = "!="

	// GreaterThan is the SQL operator for greater than
	GreaterThan SQLOperator = ">"

	// GreaterThanOrEqual is the SQL operator for greater than or equal to
	GreaterThanOrEqual SQLOperator = ">="

	// LessThan is the SQL operator for less than
	LessThan SQLOperator = "<"

	// LessThanOrEqual is the SQL operator for less than or equal to
	LessThanOrEqual SQLOperator = "<="

	// Like is the SQL operator for the LIKE operator
	Like SQLOperator = "LIKE"

	// NotLike is the SQL operator for the NOT LIKE operator
	NotLike SQLOperator = "NOT LIKE"

	// ILike is the SQL operator for the ILIKE operator
	ILike SQLOperator = "ILIKE"

	// In is the SQL operator for the IN operator
	In SQLOperator = "IN"

	// NotIn is the SQL operator for the NOT IN operator
	NotIn SQLOperator = "NOT IN"

	// IsNull is the SQL operator for the IS NULL operator
	IsNull SQLOperator = "IS NULL"

	// IsNotNull is the SQL operator for the IS NOT NULL operator
	IsNotNull SQLOperator = "IS NOT NULL"

	// Between is the SQL operator for the BETWEEN operator
	Between SQLOperator = "BETWEEN"

	// NotBetween is the SQL operator for the NOT BETWEEN operator
	NotBetween SQLOperator = "NOT BETWEEN"

	// And is the SQL operator for the AND operator
	And SQLOperator = "AND"

	// Or is the SQL operator for the OR operator
	Or SQLOperator = "OR"

	// Not is the SQL operator for the NOT operator
	Not SQLOperator = "NOT"

	// Comma is the SQL operator for the comma
	Comma SQLOperator = ","

	// AS is the SQL operator for AS operator
	AS SQLOperator = "AS"

	// RETURNING is the SQL operator for RETURNING operator
	RETURNING SQLOperator = "RETURNING"

	// SET is the SQL operator for SET operator
	SET SQLOperator = "SET"

	// WHERE is the SQL operator for WHERE operator
	WHERE SQLOperator = "WHERE"
)
