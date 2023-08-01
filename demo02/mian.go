package main

import "fmt"

/*const ( //定义多个常量,枚举类型
	BEIJING  = 1
	SHANGHAI = 2
	BAISE    = 3
)*/

const ( //定义多个常量,枚举类型
	BEIJING = iota
	SHANGHAI
	BAISE
)

func main() {
	//var 关键字
	//方法1 声明不赋值是,默认为0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("a的变量类型是:%T\n", a)

	//方法2  声明变量且赋值
	var b int = 90
	fmt.Printf("b=%d,b的类型是%T\n", b, b)

	//方法3 声明变量的时候,去掉数据类型,但是要赋值(GO语言通过值来自动匹配类型)
	var c = 100
	fmt.Printf("c=%d,c的数据类型是%T\n", c, c)

	var cc = "风萧萧兮易水寒,壮士一去兮不复还"
	fmt.Printf("%s", cc)

	//方法4 短声明(推荐)

	d := "短声明方式"
	fmt.Println(d)

	//一次声明多个变量

	//同类型的
	var a1, a2 int = 100, 300
	//不同类型的
	var s1, s2 = 123, "字符串变量"
	//其他方式
	var (
		oo    int = 999
		email     = "yv3541@qq.com"
	)

	fmt.Println(a1, a2, s1, s2, oo, email)

	//定义常量 使用const关键字
	const count int = 30
	fmt.Println(count)
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("BAISE=", BAISE)

	//iota 常量计数器

}
