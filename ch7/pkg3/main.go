package main

import (
	"fmt"
	"reflect"
)

func main1() {
	x := 1
	y := 1.1
	z := "one"
	fmt.Printf("x: %v(%v)\n", reflect.ValueOf(x).Int(), reflect.TypeOf(x))
	fmt.Printf("y: %v(%v)\n", reflect.ValueOf(y).Float(), reflect.TypeOf(y))
	fmt.Printf("z: %v(%v)\n", reflect.ValueOf(z).String(), reflect.TypeOf(z))
}

func main2() {
	type User struct {
		Name string "check:len(3,40)"
		Id   int    "check:range(1,999999)"
	}
	u := User{"Jang", 1}
	uType := reflect.TypeOf(u)
	if fName, ok := uType.FieldByName("Name"); ok {
		fmt.Println(fName.Type, fName.Name, fName.Tag)
	}
	if fId, ok := uType.FieldByName("Id"); ok {
		fmt.Println(fId.Type, fId.Name, fId.Tag)
	}
}

func main3() {
	languages := []string{"golang", "java", "c++"}
	sliceValue := reflect.ValueOf(languages)
	value := sliceValue.Index(1)
	value.SetString("ruby")
	fmt.Println(languages)
	x := 1
	if v := reflect.ValueOf(x); v.CanSet() {
		v.SetInt(2)
	}
	fmt.Println(x)
	v := reflect.ValueOf(&x)
	p := v.Elem()
	p.SetInt(3)
	fmt.Println(x)
}

func main() {
	main1()
	main2()
	main3()
}
