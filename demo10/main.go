package main

import ( //这里导入了两个包
	"fmt"
	"testGO/demo10/book"

	//给包起别名
	pp "testGO/demo10/book"

	//使用 . 就可以不需要包名直接调用
	. "testGO/demo10/book"

	// 包的匿名导入  只是执行包里的 init 函数 而不能真正被调用
	_ "fmt"
)

func init() {
	//这个函数先于 main函数执行
	fmt.Println(" main.go 中的init函数")
}

func main() {
	//原始调用
	info, err := book.ShowBookInfo("<滕王阁序>", "王勃")
	fmt.Println(info, err)

	//使用别名来进行调用
	info1, err1 := pp.ShowBookInfo("<滕王阁序>", "王勃")
	fmt.Println(info1, err1)

	//不使用报名或别名, 直接调用
	info2, err2 := ShowBookInfo("<滕王阁序>", "王勃")
	fmt.Println(info2, err2)

}
