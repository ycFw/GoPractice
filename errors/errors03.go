package main

import (
	"fmt"
	"log"
	"math"
)

func main01() {

	result, err := square(-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Square of 16 is %f", result)
}

//忽略错误
func main() {

	result, _ := square(-1)
	fmt.Printf("Square of 16 is %v\n", result)

}
func square(value float64) (float64, error) {
	if value < 0 {
		return 0, fmt.Errorf("You can not use negative numbers!")
	}
	return math.Sqrt(value), nil
}
