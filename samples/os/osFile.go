package main

import (
	"fmt"
	"io"
	"os"
)

//os.File表示本地系统上的文件，他实现了io.Reader和io.Writer接口
//展示如何将连续的字符串切片直接写入文件

func main01() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	file, err := os.Create("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range proverbs {
		//file类型实现了io.Writer
		n, err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
}

//io.File 也可以用作读取器来从本地文件系统读取文件的内容
func main() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p := make([]byte, 4)
	for {
		n, err := file.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		fmt.Print(string(p[:n]))
	}
}
