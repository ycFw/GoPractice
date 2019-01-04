package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//该函数可以方便的将字符串写入一个Writer
	file, err := os.Create("./magic_msg.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.WriteString(file, "Go is fun!"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
