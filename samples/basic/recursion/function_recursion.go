//函数式变成实现递归

package main

import "fmt"

//实现一个reserve函数
func reserve(str string) string {
	if str == "" {
		return str
	} else {
		return reserve(str[1:]) + string(str[0])
	}
}

//实现isPalindrome
// func isPalindrome(str string) string {
// 	return reserve(str)
// }

// func main01() {
// 	fmt.Println(isPalindrome("yincong"))
// }

func main() {
	isPalindrome := func(str string) string {
		return reserve(str)
	}
	fmt.Println(isPalindrome("yincong"))
}
