package main

import (
	"errors"
	"fmt"
	"math"
)

func newError() {
	errNotFound := errors.New("not found error")
	fmt.Println("error: ", errNotFound)
	fmt.Println("error: ", errNotFound.Error())
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("음수는 사용할 수 없습니다")
		return 0, fmt.Errorf("음수(%g)는 사용할 수 없습니다", f)
	}
	return math.Sqrt(f), nil
}

func mainSqrt() {
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(f)
	}
}

func main() {
	newError()
	mainSqrt()
}
