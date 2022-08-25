package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		num := fibonacciMethodOne(i)

		fmt.Printf("num: %s\n", num)
	}
}

/*
*
方法一：递归实现
*/
func fibonacciMethodOne(n int) int {
	if n < 2 {
		return 1
	}
	i := fibonacciMethodOne(n-1) + fibonacciMethodOne(n-2)
	return i
}

func fibonacciMethodTwo(n int) int {

	return 0
}
