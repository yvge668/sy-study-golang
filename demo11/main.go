package main

import "fmt"

// 流程控制
func main() {
	/*if语法*/
	score := 88
	if score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}

	//多分支判断
	score1 := 88
	if score1 >= 90 {
		fmt.Println("成绩等级为A")
	} else if score1 >= 80 {
		fmt.Println("成绩等级为B")
	} else if score1 >= 70 {
		fmt.Println("成绩等级为C")
	} else if score1 >= 60 {
		fmt.Println("成绩等级为D")
	} else {
		fmt.Println("成绩等级为E 成绩不及格")
	}

	//if的高级语法
	if score := 88; score >= 60 { //条件之前可以加一条语句,然后用分号结尾,再接条件 语句里创建的变量仅仅在if范围里生效
		fmt.Println("成绩及格")
	}

	/*switch语法*/
	grade := "B"
	switch grade { //grade作为条件,与case进行匹配,没有break
	case "A":
		fmt.Println("Your score is between 90 and 100.")
	case "B":
		fmt.Println("Your score is between 80 and 90.")
	case "C":
		fmt.Println("Your score is between 70 and 80.")
	case "D":
		fmt.Println("Your score is between 60 and 70.")
	default:
		fmt.Println("Your score is below 60.")
	}

	//一个case 里包含多个条件
	month := 5
	switch month {
	case 1, 3, 5, 7, 8, 10, 12: //多个条件之间是 或 的关系，用逗号 , 相隔
		fmt.Println("该月份有 31 天")
	case 4, 6, 9, 11:
		fmt.Println("该月份有 30 天")
	case 2:
		fmt.Println("该月份闰年为 29 天，非闰年为 28 天")
	default:
		fmt.Println("输入有误！")
	}

	//switch 的高级用法
	switch month := 5; month { //可以将上面的switch 改成这样,在表达式之前执行一条语句 month的作用于仅仅在这个switch中生效
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("该月份有 31 天")
	case 4, 6, 9, 11:
		fmt.Println("该月份有 30 天")
	case 2:
		fmt.Println("该月份闰年为 29 天，非闰年为 28 天")
	default:
		fmt.Println("输入有误！")
	}

	//switch 后可接函数
	chinese := 88
	math := 90
	english := 95

	switch getResult(chinese, math, english) {
	case true:
		fmt.Println("考试通过")
	case false:
		fmt.Println("考试未通过")
	}

	//无表达式的switch
	score2 := 88
	switch { //switch 后面的表达式是可选的。如果省略该表达式，则表示这个 switch 语句等同于 switch true
	case score2 >= 90 && score2 <= 100: //这里case 里面可以使条件表达式
		fmt.Println("grade A")
	case score2 >= 80 && score2 < 90:
		fmt.Println("grade B")
	case score2 >= 70 && score2 < 80:
		fmt.Println("grade C")
	case score2 >= 60 && score2 < 70:
		fmt.Println("grade D")
	case score2 < 60:
		fmt.Println("grade E")
	}

	//fallthrough 语句  贯穿(相当于c语言中不加break)

	s := "圣斗士星矢"
	switch {
	case s == "圣斗士星矢":
		fmt.Println("圣斗士星矢")
		fallthrough //执行下一个case
	case s == "名侦探柯南":
		fmt.Println("名侦探柯南")
	case s != "海贼王":
		fmt.Println("海贼王")
	}

	/*循环语句 go语言中只有for语句*/
	//一个表达式
	num := 0
	for num < 4 {
		fmt.Println(num)
		num++
	}

	//三个表达式
	for num := 0; num < 4; num++ { //这里num的作用域仅在for中有效
		fmt.Println(num)
	}

	//for range 表达式 迭代的时候用
	str := "大丈夫生于天地间,岂能郁郁久居人下"
	for index, value := range str { //这里range 返回两个变量,都要接收才行, 实在用不到也要用 "_" 在接收
		fmt.Printf("index %d, value %c\n", index, value)
	}

	//死循环
	/*for {
		fmt.Println("死循环!!")
	}*/

	//break 接触当前循环

	//continue 结束本次循环,开启下一个循环
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue //如果为偶数则开启下个循环,不执行下面的打印语句
		}
		fmt.Println(num)
	}

	/*defer 延迟调用*/
	defer bookPrint() //
	fmt.Println("我先执行! 才到defer 的bookPrint")
	fmt.Println("-----------------------")

	/*
						str := "LK-99复现成功!"
						defer fmt.Println(str)
						str = "阿西吧!"
						fmt.Println(str)


		即时求值的变量快照
		使用 defer 只是延时调用函数，传递给函数里的变量，不应该受到后续程序的影响。

		defer 不仅能够延迟函数的执行，也能延迟方法的执行。

			当一个函数内多次调用 defer 时，Go 会把 defer 调用放入到一个栈中，随后按照 后进先出 的顺序执行。
	*/

	showLesson()
	fmt.Println(s1)

	/*goto 语法  很少用得到~~~*/
	goto label
	fmt.Println("被无条件跳过的语句")
label:
	fmt.Println("goto到的语句")

}

func getResult(args ...int) bool { //多参数的函数
	for _, v := range args {
		if v < 60 {
			return false
		}
	}
	return true
}

func bookPrint() {
	fmt.Println("defer 我是后来执行的")
}

var s1 string = "文心一言"

func showLesson() string {
	defer func() {
		s1 = "chatGPT"
	}()
	fmt.Println("showLesson: s1 =", s1)
	return s1
}
