package database

import (
	"github.com/blacheinc/pixel/enum"
	"github.com/blacheinc/pixel/types"
)

// MapToQuery converts an types.SQLMap to an SQL query string and a slice of arguments
// suitable for use with bun.NewRaw
// if noargs is set to true, the returned slice of arguments will be empty as the value placeholders will be replaced with the actual values
func MapToQuery(m types.SQLMap, noargs ...bool) (string, []interface{}) {
	var query string
	var args []interface{}
	for k, v := range m.Map {
		if noargs != nil && noargs[0] {
			query += k + " " + string(m.ComparisonOperator) + " " + v.(string) + " " + string(m.JoinOperator) + " "
			continue
		}
		query += k + " " + string(m.ComparisonOperator) + " ? " + string(m.JoinOperator) + " "
		args = append(args, v)
	}
	query = query[:len(query)-len(string(m.JoinOperator))-2]
	return query, args
}

// MapsToWQuery converts an types.SQLMaps to an SQL read query string and a slice of arguments
// suitable for use with bun.NewRaw
func MapsToWQuery(m types.SQLMaps) (string, []interface{}) {
	var query string
	var args []interface{}
	for _, v := range m.WMaps {
		if len(v.Map) != 0 {
			q, a := MapToQuery(v)
			query += "(" + q + ") " + string(m.WJoinOperator) + " "
			args = append(args, a...)
		}
	}
	if query != "" {
		query = query[:len(query)-len(string(m.WJoinOperator))-2]
	}
	return query, args
}

// MapsToJQuery converts an types.SQLMaps to an SQL join query string and a slice of arguments
// suitable for use with bun.NewRaw
func MapsToJQuery(m types.SQLMaps) (string, []interface{}) {
	var query string
	var args []interface{}
	for _, v := range m.JMaps {
		if len(v.Map) != 0 {
			q, a := MapToQuery(v, true)
			query += "(" + q + ") " + string(m.JJoinOperator) + " "
			args = append(args, a...)
		}
	}
	if query != "" {
		query = query[:len(query)-len(string(m.JJoinOperator))-2]
	}
	return query, args
}

// MapsToSQuery converts an types.SQLMaps to an SQL update query string and a slice of arguments
// suitable for use with bun.NewRaw
func MapsToSQuery(m types.SQLMaps) (string, []interface{}) {
	wquery, wargs := MapsToWQuery(m)
	squery, sargs := MapToQuery(m.SMap)
	rquery := MapToRQuery(m.RMap)
	if rquery == "" {
		return string(enum.SET) + " " + squery + " " + string(enum.WHERE) + " " + wquery, append(sargs, wargs...)
	}
	return string(enum.SET) + " " + squery + " " + string(enum.WHERE) + " " + wquery + " " + string(enum.RETURNING) + " " + rquery, append(sargs, wargs...)
}

// MapToRQuery converts an types.SQLMap to an SQL returning query string and a slice of arguments
// suitable for use with bun.NewRaw
func MapToRQuery(m types.SQLMap) string {
	var query string
	var i int
	for k, v := range m.Map {
		if v != nil && m.ComparisonOperator != "" {
			query += k + " " + string(m.ComparisonOperator) + " " + v.(string)
			// if this is the last element, don't add a comma
			if i != len(m.Map)-1 {
				query += ", "
			}
			continue
		}
		query += k
		// if this is the last element, don't add a comma
		if i != len(m.Map)-1 {
			query += ", "
		}
		i++
	}
	return query
}
