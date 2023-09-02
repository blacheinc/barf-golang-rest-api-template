package types

import "github.com/blacheinc/pixel/enum"

type SQLMap struct {
	Map                map[string]interface{}
	JoinOperator       enum.SQLOperator
	ComparisonOperator enum.SQLOperator
}

type SQLMaps struct {
	// WMaps for WHERE clauses
	WMaps []SQLMap
	// SMaps for SET clauses
	SMap SQLMap
	// RMMap for RETURNING clause
	RMap SQLMap
	// JMaps for JOIN clauses
	JMaps []SQLMap
	// WJoinOperator for the SQLMaps present in the WMaps slice
	WJoinOperator enum.SQLOperator
	// JJoinOperator for the SQLMaps present in the JMaps slice
	JJoinOperator enum.SQLOperator
}
