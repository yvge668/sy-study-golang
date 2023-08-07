package main

import (
	"fmt"
	"time"
)

/*协程*/

func main() {
	go PrintInfo(1)             //开启携程
	go PrintInfo(2)             //开启携程
	go PrintInfo(3)             //开启携程
	time.Sleep(2 * time.Second) //睡眠两秒

}

// PrintInfo 创建一个函数,等一下开辟协程进行 调用
func PrintInfo(flag int) {
	fmt.Println("协程已开启......")
	for i := 0; i < 3; i++ {
		fmt.Println("现在运行的是协程", flag)
	}
}
