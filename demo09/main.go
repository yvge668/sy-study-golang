package main

import (
	"errors"
	"fmt"
)

// 函数
/*
	在参数类型前面加 ... 表示一个切片，
	用来接收调用者传入的参数。注意，如果该函数下有其他类型的参数，
	这些其他参数必须放在参数列表的前面，切片必须放在最后。

*/
// 多个类型一致的参数
func show(args ...string) int { //也就是说可以传入多少个字符串型的参数都行
	sum := 0
	for _, item := range args {
		fmt.Println(item)
		sum += 1
	}
	return sum // 匿名的返回值
}

// 多个类型不同的参数 使用 ...interface{}
func PrintType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "type is int.")
		case string:
			fmt.Println(arg, "type is string.")
		case float64:
			fmt.Println(arg, "type is float64.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 多个返回值的函数 string error连个类型的返回值
func showBookInfo(bookName, authorName string) (string, error) {
	if bookName == "" {
		return "", errors.New("图书名称为空")
	}
	if authorName == "" {
		return "", errors.New("作者名称为空")
	}
	return bookName + ",作者:" + authorName, nil
}

// 多个返回值, 且返回值带有变量名
func showBookInfo2(bookName, authorName string) (info string, err error) {
	info = ""
	if bookName == "" {
		err = errors.New("图书名称为空")
		return
	}
	if authorName == "" {
		err = errors.New("作者名称为空")
		return
	}
	// 不使用 := 因为已经在返回值那里声明了
	info = bookName + ",作者:" + authorName
	// 直接返回即可
	return
}

/*匿名函数
	func (parameter_list) (result_list) {
	body
}
*/

func main() {

	fmt.Printf("传入了%d个字符串\n", show("hello", "golang", "Slice", "Array"))
	PrintType(36, 3.14, "古乐天")

	/*解序列*/
	//使用 ... 可以用来解序列，能将函数的可变参数(即切片)一个一个取出来，传递给另一个可变参数的函数，而不是传递可变参数变量本身
	var s []string                                         //定义切片
	s = append(s, []string{"golang", "java", "python"}...) //追加字符串到切片中
	s = append(s, "javascript")                            //追加
	fmt.Println(s)

	bookInfo, err := showBookInfo("《三国演义》", "罗贯中")
	fmt.Printf("bookInfo = %s, err = %v\n", bookInfo, err)

	bookInfo2, err2 := showBookInfo2("《水浒传》", "施耐庵")
	fmt.Printf("bookInfo2 = %s, err2 = %v\n", bookInfo2, err2)
	/*
	   	在 Go 语言中，函数名通过首字母大小写实现控制对方法的访问权限。

	      当方法的首字母为 大写 时，这个方法对于 所有包 都是 Public ，其他包可以随意调用。
	      当方法的首字母为 小写 时，这个方法是 Private ，其他包是无法访问的。

	*/
}
