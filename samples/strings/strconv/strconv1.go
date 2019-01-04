package main

import (
	"fmt"
	"strconv"
)

//从字符串中解析数字
func main() {

	//解析浮点数,64表示解析的数的位数
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	//0表示自动推断字符串表示的数字的进制，64表示返回整数以64位存储
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	//自动识别出16进制数
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	//ParseUint也是可用的
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//Atoi是一个基础的10进制整型数转换函数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	//返回错误
	_, e := strconv.Atoi("what")
	fmt.Println(e)
}
