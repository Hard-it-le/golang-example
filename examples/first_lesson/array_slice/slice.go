package main

import "fmt"

/**
数据和切片的区别
			数组     	切片
直接初始化	支持			支持
make		不支持		支持
访问元素  	arr[i] 		arr[i]
len			长度			已有元素个数
cap			长度			容量
append		不支持		支持
是否可以扩容	不可以   	可以
*/

/**
切片

切片,语法:[]type

1.直接初始化

2. make初始化:make([]type, length, capacity)

3. arr[i]的形式访问元素

4. append 追加元素

5. len 获取元素数量

6.cap 获取切片容容量

7. 推荐写法:s1:= make([]type. 0. capacity)

如何理解切片

1。 切片操作是有限的，不支持随机增州（即没有 add， delete 方法，需要自己写代码）

2。 只有 append 操作

3。 切片支持子切片操作，和原本切片是共享底层数组
*/

func main() {
	s1 := []int{1, 2, 3, 4} // 直接初始化了 4 个元素的切片
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) // 创建了一个包含三个元素，容量为4的切片
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// s2 目前 [0, 0, 0], append（追加）一个元素，变成什么？
	s2 = append(s2, 7) // 后边添加一个元素，没有超出容量限制，不会发生扩容
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 8) // 后边添加了一个元素，触发扩容
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) // 只传入一个参数，表示创建一个含有四个元素，容量也为四个元素的
	// 等价于 s3 := make([]int, 4, 4)
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	// 按下标索引
	fmt.Printf("s3[2]: %d", s3[2])
	// 超出下标范围，直接崩溃
	// runtime error: index out of range [99] with length 4
	// fmt.Printf("s3[99]: %d", s3[99])

	// SubSlice()

	ShareSlice()
}

// SubSlice /**
/**
子切片
数组和切片都可以通过[start：end ]的形式来获取子切片：

1。 arr[start：end]，获得[start， end)之闻的元素

2。 arr[：end]，获得[O， end) 之间的元素

3。 arr[start:]，获得[start， len(arr)之间的元素
*/
func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4: %v, len %d, cap: %d \n", s4, len(s4), cap(s4))
}

// ShareSlice	/**
/**
共享底层的含义就是共享底层数组
核心：共享数组

子切片和切片究竟会不会互相影响，就抓住一点：它们是不是还共享数组？
就是如果它们结构没有变化，那肯定是共享的；

但是结构变化了，就可能不是共享了
*/

func ShareSlice() {

	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[1] = 1999
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
}
