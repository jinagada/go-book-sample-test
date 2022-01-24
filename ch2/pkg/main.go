package main

import (
	"fmt"
	mylib "go-book-sample-test/ch2/lib"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(mylib.IsDigit('1'))
	fmt.Println(mylib.IsDigit('a'))
	//fmt.Println(lib.isSpace('\t'))
}
