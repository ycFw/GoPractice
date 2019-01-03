package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Reader struct {
	buf          []byte
	rd           io.Writer
	r, w         int
	err          error
	lastByte     int
	lastRuneSize int
}

type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

// ReadWriter 集成了 bufio.Reader 和 bufio.Writer, 实现了 io.ReadWriter 接口
type ReadWriter struct {
	*Reader
	*Writer
}

//bufio package test
func main() {

	//1、readbyte和unreadbyte
	inputReadBuf1 := strings.NewReader("123456789")
	reader1 := bufio.NewReader(inputReadBuf1)

	//读一个字节
	b1, _ := reader1.ReadByte()
	fmt.Println(string(b1))
	//Unread吐出一个字节
	reader1.UnreadByte()
	b1, _ = reader1.ReadByte()
	fmt.Println(string(b1))
	fmt.Println()

	//2、readrune和unreadrune
	//readrune读出一个utf8字符并返回编码长度
	inputReadBuf2 := strings.NewReader("中文123456789")
	reader2 := bufio.NewReader(inputReadBuf2)
	b2, size, _ := reader2.ReadRune()
	fmt.Println(string(b2), size)
	reader2.UnreadRune()
	b2, size, _ = reader2.ReadRune()
	fmt.Println(string(b2), size)
	fmt.Println()

	// 执行UnreadRune时候, 如果之前一步不是ReadRune, 那么会报错
	b22, _ := reader2.ReadByte()
	fmt.Println(string(b22))
	err2 := reader2.UnreadRune()
	if err2 != nil {
		fmt.Println("ERR")
	}
	fmt.Println()

	//3、WriteTo函数
	inputReadBuf3 := strings.NewReader("中文123456789")
	reader3 := bufio.NewReader(inputReadBuf3)
	b3 := bytes.NewBuffer(make([]byte, 0))
	reader3.WriteTo(b3)
	fmt.Println(b3)
	fmt.Println()

	//4、ReadFrom函数
	inputReadBuf4 := strings.NewReader("阿萨德的发挥看书")
	b4 := bytes.NewBuffer(make([]byte, 0))
	writer4 := bufio.NewWriter(b4)
	writer4.ReadFrom(inputReadBuf4)
	fmt.Println(b4)
	fmt.Println()

	//5、使用bufio.NewReader构造一个reader
	inputReadBuf5 := strings.NewReader("1234567890")
	reader5 := bufio.NewReader(inputReadBuf5)

	//使用bufio.NewWriter构造一个writer
	buf5 := bytes.NewBuffer(make([]byte, 0))
	writer5 := bufio.NewWriter(buf5)

	//Peek: 返回缓存的一个slice，引用缓存中前n字节数据
	b, err := reader5.Peek(5)
	if err != nil {
		fmt.Printf("Read data error")
		return
	}
	//修改第一个字符
	b[0] = 'A'
	//重新读取
	b, _ = reader5.Peek(5)
	writer5.Write(b)
	writer5.Flush()
	fmt.Println("buf(changed): ", buf5, "\ninputReadBuf(Not Changed): ", inputReadBuf5)
	fmt.Println()

	//6、读取缓冲区中数据字节数(只有执行读才会使用到缓冲区, 否则是没有的)
	inputReadBuf6 := strings.NewReader("中文1234567890")
	reader6 := bufio.NewReader(inputReadBuf6)
	// 下面返回0, 因为还没有开始读取, 缓冲区没有数据
	fmt.Println(reader6.Buffered())
	// 下面返回strings的整体长度16(一个人中文是3长度)
	reader6.Peek(1)
	fmt.Println(reader6.Buffered())
	// 下面返回15, 因为readByte已经读取一个字节数据, 所以缓冲区还有15字节
	reader6.ReadByte()
	fmt.Println(reader6.Buffered())
	reader6.ReadRune()
	fmt.Println(reader6.Buffered()) //14
	reader6.ReadRune()
	fmt.Println(reader6.Buffered()) //13
	reader6.ReadRune()
	fmt.Println(reader6.Buffered()) //10

}
