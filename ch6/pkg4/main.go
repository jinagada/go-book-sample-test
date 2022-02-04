package main

import (
	"fmt"
	"log"
)

func protect(g func()) {
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("runtime panic: %v", err)
		}
	}()
	log.Println("start")
	g()
}

func divide(a, b int) int {
	return a / b
}

func main() {
	protect(func() {
		fmt.Println(divide(1, 0))
	})
}
