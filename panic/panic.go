package main

import "os"

func main() {
	panic("a problem")
	//创建新文件返回异常错误
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
