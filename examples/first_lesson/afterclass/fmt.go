package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	const f1 = 12.1123
	with2 := printNumWith2(f1)
	fmt.Printf("with2 %v:\n", with2)
	fmt.Printf("================== \n")

	const test string = "12312312"
	decodeString, _ := hex.DecodeString(test)

	fmt.Printf("string  byte : %s \n", decodeString)
	bytes := printBytes(decodeString)

	fmt.Printf("bytes : %s \n\n", bytes)
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float642), 64)
	fmt.Println(value)
	sprintf := fmt.Sprintf("value : %s \n", value)
	return sprintf
}

// 将[]byte  输出为16进制
func printBytes(data []byte) string {
	toString := hex.EncodeToString(data)
	return toString
}
