package main

const internal = "包内可访问"
const External = "包外可访问"

// 首字符是否大写控制了访问性：大写包外可访问；
// 驼峰命名
// 支持类型推断
// 无法修改值

func main() {
	const a = "你好"
	println(a)
}
