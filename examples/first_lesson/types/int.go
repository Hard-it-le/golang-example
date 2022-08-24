package main

import "fmt"

/*
*
golang 的数字类型明确标注了长度、有无符号
golang 不会帮你做类型转换，类型不同无法通过编译。也因此，string 只能和string 拼接
golang 有一个很特殊的 rune 类型， 接近一般语言的 char 或者 character 的概念，非面试情况下，可以理解为“rune=字符”
string 遇事不决找 strings 包
*/
func main() {
	// int int8、int16、int32、int64
	a := 12
	fmt.Println(a)

	// uint uint8、uint16、uint32、uint64、uint
	b := 23
	fmt.Println(b)
	// bool true、bool false
	c := true
	for i := 0; i < a; i++ {
		c = false
	}
	// float32、float64
	fmt.Println(c)
}
