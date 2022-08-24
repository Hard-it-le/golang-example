package main

// 局部变量，只适合包内变量,golang 使用类型推断来推断类型。数字会被理解为 int 或者 float64。（所以要其它类型的数字，就得用 var 来声明）

func main() {
	a := 13
	println(a)
	b := "你好"
	println(b)
}
