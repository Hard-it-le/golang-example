package main

import "fmt"

/** fun 方法声明 结构
	func  方法名字()[参数列表][返回列表]{}
	方法名字：首字母是否大写决定了作用域

	推荐写法
	- 参数列表含有参数名
	- 返回值不具有返回值名
如 Fun0 和 Fun1
*/
/**
方法声明和方法调用总结
golang 支持多返回值，这是一个很大的不同点

golang 方法的作用域和变量作用域一样，通过大小写控制

golang 的返回值是可以有名字的，可以通过给予名字让调用方清楚知道你返回的是什么
*/

func main() {
	a := Fun0("Tom")
	println(a)

	b, c := Fun1("a", 17)
	println(b)
	println(c)

	// 方法的调用 使用_忽略返回值
	_, d := Fun2("a", "b")
	println(d)

	// 不定参数后面可以传递任意多个值
	Fun4("hello", 19, "CUICUI", "DaMing")
	s := []string{"CUICUI", "DaMing"}
	Fun4("hello", 19, s...)
}

// Fun0 只有一个返回值，不需要括号括起来
func Fun0(name string) string {
	return "Hello, " + name
}

// Fun1 多个参数，多个返回值。参数有名字，但是返回值没有
func Fun1(a string, b int) (int, string) {
	return 0, "你好"
}

// Fun2 的返回值具有名字，可以在内部直接复制，然后返回
// 也可以忽略age, name，直接返回别的。
func Fun2(a string, b string) (age int, name string) {
	age = 19
	name = "Tom"
	return
	//return 19, "Tom" // 这样返回也可以
}

// Fun3 多个参数具有相同类型放在一起，可以只写一次类型
func Fun3(a, b, c string, abc, bcd int, p string) (d, e int, g string) {
	d = 15
	e = 16
	g = "你好"
	return
	//return 0, 0, "你好" // 这样也可以
}

// Fun4 不定参数。不定参数要放在最后面
func Fun4(a string, b int, names ...string) {
	// 我们使用的时候可以直接把 names 看做切片
	for _, name := range names {
		fmt.Printf("不定参数：%s \n", name)
	}
}
