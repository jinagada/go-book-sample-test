package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	return a / b
}

func main1() {
	fmt.Println(divide(1, 0))
}

func main2() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

func main() {
	main1()
	//main2()
}
