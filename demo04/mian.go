package main

import "fmt"

func test01() {
	// 声明时没有指定数组元素的值, 默认为零值
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
}

func test02() {
	// 直接在声明时对数组进行初始化
	var arr1 = [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr1)

	// 使用短声明
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr2)

	// 部分初始化, 未初始化的为零值
	arr3 := [5]int{15, 20} // [15 20 0 0 0]
	fmt.Println(arr3)

	// 可以通过指定索引，方便地对数组某几个元素赋值
	arr4 := [5]int{1: 100, 4: 200}
	fmt.Println(arr4) // [0 100 0 0 200]

	// 直接使用 ... 让编译器为我们计算该数组的长度
	arr5 := [...]int{15, 20, 25, 30, 35, 40}
	fmt.Println(arr5)
}

func test03() {
	// 特别注意数组的长度是类型的一部分，所以 [3]int 和 [5]int 是不同的类型
	arr1 := [3]int{15, 20, 25}
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Printf("type of arr1 is %T\n", arr1)
	fmt.Printf("type of arr2 is %T\n", arr2)
}

func test04() {
	// 定义多维数组
	arr := [3][2]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"}}
	fmt.Println(arr) // [[15 20] [25 22] [25 22]]
}

func showArr() {
	//数组遍历
	arr := []string{"java", "golang", "php", "javaScript", "python"}

	//使用 for range 必须要用 键和值来接,必须要有两个
	for index, value := range arr {
		fmt.Printf("arr[%d]=%s\n", index, value)
	}

	for _, value := range arr {
		fmt.Println("value=", value)
	}
}

func arrByValue() {
	//go语言中,数组是值类型(基本数据类型),不是引用类型,当数组赋值给一个新的变量时,该变量会得到一个原始数组的一个副本.如果对新变量进行更改,不会影响原始数组
	arr := [...]string{"for i", "for range", "switch"}
	copy1 := arr
	copy1[0] = "宇宙超级无敌QQ哥"

	fmt.Println(arr)
	fmt.Println(copy1)

}
func main() {
	/*
		数组是值类型 不是引用数据类型
	*/

	test01()
	test02()
	test03()
	test04()
	fmt.Println()
	showArr()
	fmt.Println()
	arrByValue()
}
