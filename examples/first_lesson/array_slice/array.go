package main

import "fmt"

/*
*
数组和别的语言的数组差不多，语法是：[cap]type

1。 初始化要指定长度（或者叫做容量）

2。直接初始化

3。 arr[i]的形式访问元素

4。 len 和 cap 操作用于获取数组长度
*/
func main() {
	// 直接初始化一个三个元素的数组。大括号里面多一个或者少一个都编译不通过
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: %v, len: %d, cap: %d \n", a1, len(a1), cap(a1))

	// 初始化一个三个元素的数组，所有元素都是0
	var a2 [3]int
	fmt.Printf("a2: %v, len: %d, cap: %d \n", a2, len(a2), cap(a2))

	//a1 = append(a1, 12) 数组不支持 append 操作

	// 按下标索引
	fmt.Printf("a1[1]: %d", a1[1])
	// 超出下标范围，直接崩溃，编译不通过
	//fmt.Printf("a1[99]: %d", a1[99])
}
