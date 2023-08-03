package book

import (
	"errors"
	"fmt"
)

func init() {
	//当加载这个包的时后就会执行init函数
	fmt.Println(" book.go 中的init函数")
}

// 函数名必须使用大写开头  不然外层的main访问不到这个包下的这个函数
func ShowBookInfo(bookName, authorName string) (string, error) {
	if bookName == "" {
		return "", errors.New("图书名称为空")
	}
	if authorName == "" {
		return "", errors.New("作者名称为空")
	}
	return bookName + ",作者:" + authorName, nil
}
