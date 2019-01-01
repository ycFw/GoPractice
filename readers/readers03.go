package main

import (
	"fmt"
	"io"
)

//自己实现一个reader，而不是用标准库的io.reader读取器
//实现从流中过滤掉非字母字符的功能
type alphaReader struct {
	src string //资源
	cur int    //读取到当前位置
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

//过滤函数
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

//Read方法
func (a *alphaReader) Read(p []byte) (int, error) {
	//当前位置 >= 字符串长度，说明已经读取到结尾，返回 EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}
	//x是剩余未读取的长度
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		//剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		//剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		//每次读取一个字节，执行过滤函数
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	//将处理后得到的buf内容复制到p中
	copy(p, buf)
	return n, nil

}

func main() {
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			// fmt.Println("EOF: ", n)
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
