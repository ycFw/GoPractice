package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v ---> %s", e.When, e.What)
}

//自定义error发生时，显示的输出格式
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main01() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

//Sqrt 函数接收负数返回error
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return 0, nil
	} else {
		return x, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

}
