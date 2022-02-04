package main

import (
	"fmt"
	"reflect"
)

func Len(x interface{}) int {
	value := reflect.ValueOf(x)
	switch reflect.TypeOf(x).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return value.Len()
	default:
		if method := value.MethodByName("Len"); method.IsValid() {
			values := method.Call(nil)
			return int(values[0].Int())
		}
	}
	panic(fmt.Sprintf("'%v' does not have a length", x))
}
