package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	//每次以8个字节读取strings.NewReader的输出
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8) //8  这里控制每次读取的字节数
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v		err = %v	b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
