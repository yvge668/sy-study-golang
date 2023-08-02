package main

import "fmt"

// 指针
func main() {
	//通过定义普通变量获取指针

	x := "定义了一个字符串"
	ptr := &x
	fmt.Println(ptr) //打印地址
	fmt.Println(*ptr)

	ptr2 := new(string)
	*ptr2 = "new先创建指针,然后现在才是给指针指向的地址赋一个字符串"
	fmt.Println(ptr2)
	fmt.Println(*ptr2)
	fmt.Println()
	x2 := "这是一个字符串变量"
	var p *string
	p11 := &x2
	p = &x2

	fmt.Println(p11)
	fmt.Println(*p11)
	fmt.Println(p)
	fmt.Println(*p)

	//& 取地址
	//* 取值

}

func pointerType() {
	//*(指向变量值的数据类型) 就是对应的指针类型。

	mystr := "字符串类型"
	myint := 1
	mybool := false
	myfloat := 3.2
	fmt.Printf("type of &mystr is :%T\n", &mystr)
	fmt.Printf("type of &myint is :%T\n", &myint)
	fmt.Printf("type of &mybool is :%T\n", &mybool)
	fmt.Printf("type of &myfloat is :%T\n", &myfloat)
}

func zeroPointer() {
	//指针的默认值

	x := "我是一个字符串"
	var ptr *string             //创建指针 但为赋地址
	fmt.Println("ptr is ", ptr) //默认是nil
	ptr = &x
	fmt.Println("ptr is ", ptr)
}

//切片与指针一样是引用类型，如果我们想通过一个函数改变一个数组的值，可以将该数组的切片当作参数传给函数，
//也可以将这个数组的指针当作参数传给函数。但 Go 中建议使用第一种方法，
//即将该数组的切片当作参数传给函数，因为这么写更加简洁易读

// 使用切片 (引用类型 值就是一个地址)
func changeSlice(value []int) { //这里不需要* . 因为切片本身就是引用类型的数据(地址)
	value[0] = 200
}

// 使用数组指针 数组是一个值类型的数据 这里的意思是要传进一个value的地址
func changeArray(value *[3]int) {
	(*value)[0] = 200
}
