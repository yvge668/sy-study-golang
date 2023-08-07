package main

import "fmt"

// 定义动物接口
type Animal interface {
	run() //跑的接口方法
}

// 定义结构体
type Dog struct { //狗子类型结构体
	name string
	age  int
}

type Cat struct { //小猫类型结构体
	name  string
	age   int
	skill string
}

// 给结构体绑定接口方法
func (d Dog) run() {
	fmt.Printf("%d岁的狗子%s正在飞快的跑\n", d.age, d.name)
}

func (c *Cat) run() {
	fmt.Printf("%d岁的小猫%s正在%s\n", c.age, c.name, c.skill)
}

func main() {

	dog1 := Dog{
		name: "旺财",
		age:  2,
	}
	dog1.run()

	cat1 := Cat{
		name:  "花花",
		age:   1,
		skill: "睡觉",
	}
	cat1.run()

	dog2 := Dog{
		name: "富贵",
		age:  1,
	}

	dog3 := dog2
	dog3.run()

	cat2 := Cat{
		name:  "朵朵",
		age:   1,
		skill: "撒娇",
	}

	cat3 := &cat2
	cat3.run()

	dog4 := Dog{
		name: "常威",
		age:  1,
	}

	dog5 := dog4
	dog6 := &dog4
	//不太理解的是这里, 明明两个都不一样,一个是变量,一个是指针,为啥都可以调用run方法呀?
	//我这里猜测,应该是应该是类似于c语言的 dog5.run  和 dog6->run
	fmt.Printf("%T\n", dog5)
	fmt.Printf("%T\n", dog6)
	dog5.run()
	dog6.run()

	cat4 := Cat{
		name:  "丹丹",
		age:   1,
		skill: "撒娇",
	}
	cat5 := &cat4 //注意现在cat5是一个指针

	cat5.run() //用指针也可以调用

	fmt.Printf("--------------------------------------\n")

	//定义两个接口变量
	var a1, a2 Animal

	a1 = dog5 //小狗结构体变量赋值给了 接口变量
	a2 = cat5 //小猫结构体指针赋值给了接口变量

	a1.run() //接口变量调用
	a2.run() //接口指针调用

	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Printf("%x\n", &a1)
	fmt.Printf("%x\n", &a2)

	ShowInterface(a1)
	fmt.Println("=====================")
	ShowInterface(a2)

}

func ShowInterface(a Animal) {
	fmt.Printf("接口类型: %T\n,接口值: %v\n", a, a)
}

// ShowType 空接口
func ShowType(i interface{}) {
	fmt.Printf("类型: %T, 值: %v\n", i, i)
}
