package main

import "fmt"

func Sum(len int) int {
	sum := 0
	for i := 1; i <= len; i++ {
		sum += i
	}
	return sum
}
func main() {
	fmt.Println(Sum(1000))
}
