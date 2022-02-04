package main

import (
	"fmt"
	"reflect"
	"strings"
)

func TitleCase(s string) string {
	return strings.Title(s)
}

func main() {
	caption := "go is an open source programming language"
	title := TitleCase(caption)
	fmt.Println(title)
	titleFuncValue := reflect.ValueOf(TitleCase)
	values := titleFuncValue.Call([]reflect.Value{reflect.ValueOf(caption)})
	title = values[0].String()
	fmt.Println(title)
}
