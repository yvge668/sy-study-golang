package main //声明main包
import (
	"fmt"
	"os"
)

//导入fmt包, 打印用的包

func main() { //声明main
	fmt.Println("hello Go!")
	fmt.Println("2023年8月1日")
	os.Getegid()
}

/*
func 函数名(参数列表)(返回值列表){
		方法体
}
*/
