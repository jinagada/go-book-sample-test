package main

import (
	"fmt"
	"log"
)

func divide(a, b int) int {
	return a / b
}

type fType func(int, int) int

func errorHandler(fn fType) fType {
	return func(a, b int) int {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("runtime panic: %v", err)
			}
		}()
		return fn(a, b)
	}
}

func main() {
	fmt.Println(errorHandler(divide)(4, 2))
	fmt.Println(errorHandler(divide)(3, 0))
}
