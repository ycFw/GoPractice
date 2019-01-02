package main

import "fmt"
//BubbleSort
func BubbleSort(buf []int) []int {

	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 1; j < len(buf)-i; j++ {
			times ++ 
			tmp, buf[j-1] := buf[j-1], buf[j]
			flag = true
		}
	}

	if !flag{
		break
	}
	fmt.Println("BubbleSort times: ",times)
}

func main() {
	fmt.Println(BubbleSort([]int{2,5,9,10,8,1,4}))
}
