package main

import "fmt"

/*
通道(channel) ，就是一个管道，可以想像成 Go 协程之间通信的管道。它是一种队列式的数据结构，遵循先入先出的规则
*/
func main() {
	/*//声明了一个string 类型的通道
	var ch chan string //但是初值为nil

	//创建一个通道
	ch = make(chan string)*/

	//直接短声明  创建
	ch := make(chan string) //有了通道我们就可以发送数据和接收数据

	/*
			发送数据
			channel_name <- data

			接收数据
			value := <- channel

		看箭头的方向就很容易区分了
	*/

	//开启协程
	go testChannel(ch)

	//接受协程通道发过来的数据 (主程序中断等待,直到接收到管道的数据才进行下一步)
	value := <-ch

	//打印数据
	fmt.Println(value)

	//关闭协程
	value, ok := <-ch
	if ok == true { //通道还未关闭 就是true, 关闭了 是false
		close(ch)
	}

	/*
	   我们在前面讲过 make 函数是可以接收两个参数的，同理，创建通道可以传入第二个参数——容量。

	   当容量为 0 时，说明通道中不能存放数据，在发送数据时，必须要求立马有人接收，否则会报错。此时的通道称之为无缓冲通道。
	   当容量为 1 时，说明通道只能缓存一个数据，若通道中已有一个数据，此时再往里发送数据，会造成程序阻塞。利用这点可以利用通道来做锁。
	   当容量大于 1 时，通道中可以存放多个数据，可以用于多个协程之间的通信管道，共享资源。
	   既然通道有容量和长度，那么我们可以通过 cap 函数和 len 函数获取通道的容量和长度。

	*/
	// 创建一个通道
	c := make(chan int, 3)
	fmt.Println("初始化后：")
	fmt.Println("cap =", cap(c))
	fmt.Println("len =", len(c)) //现在通道里面有0个数据
	c <- 1
	c <- 2
	fmt.Println("传入两个数后：")
	fmt.Println("cap =", cap(c))
	fmt.Println("len =", len(c)) //现在通道里面有两个数据
	<-c
	fmt.Println("取出一个数后：")
	fmt.Println("cap =", cap(c))
	fmt.Println("len =", len(c)) //取出来一个数据, 现在还剩一个数据
}

func testChannel(ch chan string) {
	//向通道中发送数据(发送与接收默认程序是阻塞的, 或者说最中断等待)
	ch <- "死亡如风,常伴吾身"

}
