package main

import (
	"bytes"
	"io"
	"os"
)

//类型 io.PipeWriter 和 io.PipeReader 在内存管道中模拟 io 操作
// 数据被写入管道的一端，并使用单独的 goroutine 在管道的另一端读取
// 下面使用 io.Pipe() 创建管道的 reader 和 writer，然后将数据从 proverbs 缓冲区复制到io.Stdout

func main() {

	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	piper, pipew := io.Pipe()

	//将proverbs写入pipew这一端
	go func() {
		defer pipew.Close()
		io.Copy(pipew, proverbs)
	}()

	//从另一端piper中读取数据并拷贝到标准输出
	io.Copy(os.Stdout, piper)
	piper.Close()
}
