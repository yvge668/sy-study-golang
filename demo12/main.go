package main

import "fmt"

// 方法 注意: 方法和函数是有区别的
//在 Go 中，相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的
//每个结构体都可以有和别的结构体名字相同的方法 就像java 中 每个对象(或者说类)都可以自己的方法 即便这个方法名字和别的对象方法名字一样
//若要改变实例的属性时，必须使用指针作为方法的接收者

// Lesson 定义一个名为 Lesson 的结构体
type Lesson struct {
	Name      string
	Target    string
	SpendTime int
}
type myInt int

// 非结构体的方法
func (a myInt) add(b myInt) myInt {
	return a + b
}

// PrintInfo 定义一个与 Lesson 的绑定的方法
func (lesson Lesson) PrintInfo() { //值接收器  里面的lesson 类似于java 中的this关键字
	fmt.Println("name:", lesson.Name)
	fmt.Println("target:", lesson.Target)
	fmt.Println("spendTime:", lesson.SpendTime)
}

func (lesson Lesson) ChangeLessonName(name string) {
	lesson.Name = name
}

// AddSpendTime 定义一个与 Person 的绑定的方法，使 age 值加 n
func (lesson *Lesson) AddSpendTime(n int) { //指针接收器   只有用指针接收器 才能修改结构体的属性
	lesson.SpendTime = lesson.SpendTime + n
}

func main() {
	lesson := Lesson{
		Name:      "范冰冰",
		Target:    "国民女神~",
		SpendTime: 33,
	}
	//通过结构体调用 绑定这个结构体的方法  有点类似于java 通过对象调用对象方法
	lesson.PrintInfo()
	lesson.AddSpendTime(3)
	lesson.PrintInfo()
	fmt.Println("---------------------------------")
	//使用值参数 PrintInfo(lesson) 来调用这个函数是合法的，使用值接收器来调用 lesson.PrintInfo() 也是合法的
	PrintInfo(lesson)
	fmt.Println("-----------------------------------")

	bPtr := &lesson
	PrintInfo(*bPtr) // 把指针转成结构体变量后就可以合法使用
	fmt.Println("-----------------------------------")
	bPtr.PrintInfo()

	fmt.Println("-----------------------------------")
	var x myInt = 50
	var y myInt = 7
	fmt.Println(x.add(y)) // 57

}

func PrintInfo(lesson Lesson) { //这个是一个函数 传入的参数是 一个结构体
	fmt.Println(lesson.Name)
	fmt.Println(lesson.Target)
	fmt.Println(lesson.SpendTime)
}
