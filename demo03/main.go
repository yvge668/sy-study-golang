package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Integer 有符号类型
func Integer() {
	var num8 int8 = 127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	//unsafe.Sizeof()  获取这个变量占多少个字节, 一个字节8位
	fmt.Printf("num8的类型是:%T num8的大小是:%d num8是:%d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是:%T num16的大小是:%d num16是:%d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是:%T num32的大小是:%d num32是:%d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是:%T num64的大小是:%d num64是:%d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是:%T num的大小是:%d num是:%d\n", num, unsafe.Sizeof(num), num)
}

// 无符号整型
func unsignedInteger() {
	var num8 uint8 = 128
	var num16 uint16 = 32768
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	var num uint = math.MaxUint
	fmt.Printf("num8的类型是 %T, num8的大小 %d, num8是 %d\n",
		num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T, num16的大小 %d, num16是 %d\n",
		num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T, num32的大小 %d, num32是 %d\n",
		num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T, num64的大小 %d, num64是 %d\n",
		num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T, num的大小 %d, num是 %d\n",
		num, unsafe.Sizeof(num), num)
}
func showFloat() {
	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64
	fmt.Println("小字表示(% f)")
	fmt.Printf("num1的数据类型是%T,num1的大小是 %f\n", num1, num1)
	fmt.Printf("num1的数据类型是%T,num1的大小是 %f\n", num2, num2)
	fmt.Println("科学计数表示(% g)")
	fmt.Printf("num1的数据类型是%T,num1的大小是 %g\n", num1, num1)
	fmt.Printf("num1的数据类型是%T,num1的大小是 %g\n", num2, num2)

	fmt.Println("float32 的精度只能提供大约 6 个十进制数(表示小数点后 6 位)的精度。" +
		"\nfloat64 的精度能提供大约 15 个十进制数(表示小数点后 15 位)的精度。")
}

func main() {
	//一般使用int 表示整形的宽度,在32位系统下就是32位,64位系统下就是 64位
	Integer()
	fmt.Println()
	unsignedInteger()
	fmt.Println()
	showFloat()
}
