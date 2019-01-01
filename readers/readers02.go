package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//创建一个字符读取器，并流式的按字节读取
	readers := strings.NewReader("Clever is better than clever hahaha")
	//缓冲区p
	p := make([]byte, 4)

	for {
		n, err := readers.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF: ", n)
				break
			}
			fmt.Println(err)
			//os.exit执行，defer函数不会被执行
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}
