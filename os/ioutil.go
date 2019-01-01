package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//使用函数ioutil将文件内容加载到bytes中
	bytes, err := ioutil.ReadFile("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", bytes)
}
