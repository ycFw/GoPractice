// 使用 `os.Exit` 来立即进行带给定状态的退出

package main

import (
	"fmt"
	"os"
)

func main() {
	//永远不会执行
	defer fmt.Println("!")

	os.Exit(3)
}
