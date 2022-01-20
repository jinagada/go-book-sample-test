package main

import (
	"fmt"
	"go-book-sample-test/ch2/lib"
)

var v rune

func init() {
	v = '1'
}

func main() {
	fmt.Println(lib.IsDigit(v))
}
