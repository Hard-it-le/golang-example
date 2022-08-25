package main

import "fmt"

/*
*
带局部变量声明的 if- else：
1。 distance 只能在计块，或者后边所有的 else 块里面使用
2。脱离了f-else 块，则不能再使用
*/
func main() {
	Young(9)
	Young(100)

	IfUsingNewVariable(10, 200)
	IfUsingNewVariable(100, 30)
}

func Young(age int) {
	if age < 18 {
		fmt.Println("I am a child!")
	} else {
		// else 分支也可以没有
		fmt.Println("I not a child")
	}
}

func IfUsingNewVariable(start int, end int) {
	if distance := end - start; distance > 100 {
		fmt.Printf("距离太远，不来了： %d\n", distance)
	} else {
		// else 分支也可以没有
		fmt.Printf("距离并不远，来一趟： %d\n", distance)
	}

	// 这里不能访问  distance
	//fmt.Printf("距离是： %d\n", distance)
}
