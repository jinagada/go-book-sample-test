package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

type SqrtError struct {
	time    time.Time
	value   float64
	message string
}

func (e SqrtError) Error() string {
	return fmt.Sprintf("[%v]ERROR - %s(value: %g)", e.time, e.message, e.value)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, SqrtError{time: time.Now(), value: f, message: "음수는 사용할 수 없습니다"}
	}
	if math.IsInf(f, 1) {
		return 0, SqrtError{time: time.Now(), value: f, message: "무한대 값은 사용할 수 없습니다"}
	}
	if math.IsNaN(f) {
		return 0, SqrtError{time: time.Now(), value: f, message: "잘못된 수 입니다"}
	}
	return math.Sqrt(f), nil
}

func main() {
	v, err := Sqrt(9)
	if e, ok := err.(SqrtError); ok {
		fmt.Println("Sqrt Error", e)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	v, err = Sqrt(-9)
	switch e := err.(type) {
	case SqrtError:
		fmt.Println("Sqrt Error", e)
	default:
		fmt.Println("Default Error", e)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
