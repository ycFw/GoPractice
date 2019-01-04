package main

import (
	"fmt"
	"os"
)

//os有三个可用变量os.Stdout ，os.Stdin 和 os.Stderr
func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	for _, p := range proverbs {
		//因为 os.Stdout 也实现了 io.Writer
		//直接打印到标准输出
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
}
