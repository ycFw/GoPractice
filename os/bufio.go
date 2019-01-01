package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//bufio 包支持缓冲区io操作，可以轻松处理文本内容

func main() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	//逐行读取文件内容，并以值 '\n' 分隔
	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Print(line)
	}

}
