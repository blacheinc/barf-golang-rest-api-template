package reflection

import (
	"fmt"
	"reflect"

	"github.com/blacheinc/pixel/types"
)

/*
ReturnStructFields recursively returns
a slice of interface pointers to the fields of a struct
as need for scanning a database row into a struct
note: the order of the fields in the slice
is the same as the order of the fields in the struct
and this inturn is the same as the order of the columns in the database query

it uses the struct tag `rsf:"false"` to determine fields to skip and the tag `rsfr:"false"` to determine fields to skip on recursive calls
*/
func ReturnStructFields(s interface{}) []interface{} {
	// get the type of argument
	t := reflect.TypeOf(s)
	if t == nil {
		return nil
	}
	// only allow struct type
	if t.Elem().Kind() != reflect.Struct {
		return nil
	}
	// create a slice of interface pointers
	fields := make([]interface{}, t.Elem().NumField())
	// fill the slice with pointers to each struct field
	previousLen := 0
	for i := 0; i < t.Elem().NumField(); i++ {
		// only allow exported fields (`rsf="true"`) or recursive calls
		rsf := t.Elem().Field(i).Tag.Get("rsf") != "false"
		if rsf {
			// if the field is a struct, call this function recursively but ignore pgtype fields
			var primitives types.StringArray = []string{"pgtype", "time", "mysql"}
			if t.Elem().Field(i).Type.Kind() == reflect.Struct && (t.Elem().Field(i).Tag.Get("rsfr") != "false" && !primitives.ExistsIn(fmt.Sprintf("%T", reflect.New(t.Elem().Field(i).Type).Interface()))) {
				fields = append(fields[:previousLen], append(ReturnStructFields(reflect.ValueOf(s).Elem().Field(i).Addr().Interface()), fields[previousLen:]...)...)
				previousLen += t.Elem().Field(i).Type.NumField()
			} else {
				fields[previousLen] = reflect.ValueOf(s).Elem().Field(i).Addr().Interface()
				previousLen += 1
			}
		}
	}
	// // remove nil values
	for i := 0; i < len(fields); i++ {
		if fields[i] == nil {
			fields = append(fields[:i], fields[i+1:]...)
			i--
		}
	}
	return fields
}
