package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Print("이름을 입력하세요: ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "Hello %s\n", name)
}
