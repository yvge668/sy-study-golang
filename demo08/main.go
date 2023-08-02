package main

import (
	"fmt"
	"reflect"
)

/*创建结构体*/
//方式一
type Student struct {
	name string
	sex  string
	age  int
}

// 方式二
type Student2 struct {
	name, sex string //结构体属性同类型可以连续定义
	age       int
}

// 匿名字段 (和匿名结构体是两个概念)
type Lesson4 struct {
	string
	int
}

// 嵌套结构体
type Author struct {
	name string
	wx   string
}

type Lesson5 struct {
	name, target string
	spend        int
	author       Author
}

// 提升字段
type Lesson6 struct {
	name, target string
	spend        int
	Author
}

/*给结构体绑定方法*/
func (l Student) ShowLessonInfo() {
	fmt.Println("name:", l.name)
	fmt.Println("target:", l.sex)
	fmt.Println("target:", l.age)
}

// 如果绑定结构体的方法中要改变实例的属性时，必须使用指针作为方法的接收者
func (l *Student) editName(newName string) {
	l.name = newName
}

func main() {

	/*给结构体赋值: 通过属性名称一一赋值*/
	student := Student{
		name: "张三丰",
		sex:  "男",
		age:  66,
	}
	/*给结构体赋值: 通过属性的顺序进行赋值*/
	student1 := Student{"王天一", "男", 36}
	fmt.Println(student)
	fmt.Println(student1)

	//创建匿名结构体
	var Lesson3 struct {
		name, target string
		spend        int
	}
	//当定义好的结构体没有被显式初始化时，结构体的字段将会默认赋为相应类型的零值。
	fmt.Println(Lesson3)

	cat1 := struct {
		brand string
		price int
	}{"宝马", 100}

	cat2 := struct {
		brand string
		price int
	}{
		brand: "奔驰",
		price: 66,
	}

	fmt.Println(cat1)
	fmt.Println(cat2)
	fmt.Println("---------------------------------------")

	/*访问结构体的某一字段*/
	chessMaster := struct {
		name string
		sex  string
		age  int
	}{"许银川", "男", 35}

	/*通过指针的方式*/
	chessMaster1 := &chessMaster

	fmt.Printf("特级大师:%s已经%d了\n", chessMaster.name, chessMaster.age)
	//通过指针的方式访问, 值得注意的是 c语言中访问指针结构体时 用 -> go语言中统一用 .
	fmt.Printf("特级大师:%s已经%d了\n", (*chessMaster1).name, (*chessMaster1).age)
	fmt.Println("---------------------------------------")

	//定义一个匿名字段的结构体
	lesson4 := Lesson4{
		string: "匿名字段名称",
		int:    1,
	}

	//访问匿名字段
	fmt.Printf("当前访问string: %s\n", lesson4.string)
	fmt.Printf("当前访问int: %d\n", lesson4.int)
	fmt.Println("---------------------------------------")

	//创建嵌套结构体
	lesson10 := Lesson5{
		name:  "豹子头",
		spend: 50,
	}
	lesson10.author = Author{
		name: "林冲",
		wx:   "yvge6666",
	}
	fmt.Println("lesson10 name:", lesson10.name)
	fmt.Println("lesson10 spend:", lesson10.spend)
	fmt.Println("lesson10 author name:", lesson10.author.name)
	fmt.Println("lesson10 author wx:", lesson10.author.wx)

	fmt.Println("---------------------------------------")
	fmt.Println("---------------提升字段------------------------")
	lesson11 := Lesson6{
		name:   "外部的name1",
		target: "<{=．．．．(嘎~嘎~嘎~)",
	}
	lesson11.Author = Author{
		name: "内部的name2",
		wx:   "Author的wx属性",
	}
	fmt.Println("lesson11 name:", lesson11.name)
	fmt.Println("lesson11 target:", lesson11.target)
	fmt.Println("lesson11 author wx:", lesson11.wx)          //如果没有相同名字的属性 可以直接访问内部的属性
	fmt.Println("lesson11 author wx:", lesson11.Author.name) // 又相同的类型的属性要 多.一次
	fmt.Println("-------------结构体比较-------------------------")
	/*结构体的比较*/
	s1 := Student{
		name: "李云龙",
		sex:  "男",
		age:  40,
	}
	s2 := Student{
		name: "李云龙",
		sex:  "男",
		age:  40,
	}
	fmt.Println(s1 == s2)                  //	true  进行每个字段的比较
	fmt.Println(reflect.DeepEqual(s1, s2)) //	true 进行每个字段的比较

	/*结构体的方法*/
	sl := Student{
		name: "李小龙",
		sex:  "男",
		age:  22,
	}
	//用.的方式调用, 有点类似于java 的对象调用方法
	sl.ShowLessonInfo()
	sl.editName("李大龙")
	sl.ShowLessonInfo()

}
