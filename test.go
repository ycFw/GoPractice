package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("I love my mother.")
	p := make([]byte, 2)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF: ", n)
				break
			}
			fmt.Println("err =", err)
		}
		fmt.Println(n, string(p[:n]))
	}
}
