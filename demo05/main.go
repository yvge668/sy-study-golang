package main

import "fmt"

/*
切片是对数组的一个连续片段的引用，所以切片是一个引用类型。
切片 本身不拥有任何数据，它们只是对现有数组的引用，
每个切片值都会将数组作为其底层的数据结构。slice 的语法和数组很像，只是没有固定长度而已。
*/
func createSlice() {
	//使用 []Type 可以创建一个带有 Type 类型元素的切片。
	//声明整型切片
	var numList []int

	//声明一个空切片
	var numListEmpty = []int{}

	//我们也可以使用 make 函数构造一个切片，格式为 make([]Type, size, cap)
	numList1 := make([]string, 3, 5) //类型,当前长度,容量

	//当然，我们可以通过对数组进行片段截取创建一个切片。
	arr := [5]string{"java", "css", "JavaScript", "php", "golang"}
	var s1 = arr[1:4] //包头不包尾

	fmt.Println(numList)
	fmt.Println(numListEmpty)
	fmt.Println(numList1)
	fmt.Println(arr) //[java css JavaScript php golang]
	fmt.Println(s1)  //[css JavaScript php]

}

func SliceProperty() {
	//切片的属性
	s := make([]int, 3, 5)           //切片类型,切片长度,切片容量
	fmt.Println("当前切片的长度为:", len(s)) //3
	fmt.Println("当前切片的容量为:", cap(s)) //5

	//如果切片操作超出上限将导致一个 panic 异常。
	//fmt.Println(s[10]) //panic: runtime error: index out of range [10] with length 3

	//切片是引用类型, 所以你不对它进行赋值的话，它的默认值是 nil
	var numList []int
	fmt.Println(numList == nil) // true

	//切片之间不能比较，因此我们不能使用 == 操作符来判断两个 slice 是否含有全部相等元素。
	//特别注意，如果你需要测试一个 slice 是否是空的，使用 len(s) == 0 来判断，而不应该用 s == nil 来判断。
}

func modifySlice() {
	//修改切片元素
	//切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。
	var arr = [...]string{"modify", "Slice", "Array"}
	s := arr[:] //[0:len(arr)]
	fmt.Println(arr)
	fmt.Println(s)

	s[0] = "update" //

	fmt.Println(arr)
	fmt.Println(s)
}

func appendSliceData() {
	s := []string{"切片元素0"}
	fmt.Println(s)
	fmt.Println("当前容量", cap(s))

	s = append(s, "追加元素1")
	fmt.Println(s)
	fmt.Println("容量翻倍", cap(s))

	s = append(s, "追加元素2", "追加元素3")
	fmt.Println(s)
	fmt.Println("容量翻倍", cap(s))

	s = append(s, []string{"追加元素4", "追加元素5"}...)
	fmt.Println(s)
	fmt.Println("容量翻倍", cap(s))
}

func mSlice() { //类似于数组，切片也可以有多个维度
	numList := [][]string{
		{"1", "JavaScript"},
		{"2", "PHP"},
		{"3", "golang"},
	}
	fmt.Println(numList)
}
func main() {
	createSlice()
	SliceProperty()
	modifySlice()
	appendSliceData()
	mSlice()

	var testSlice []int
	fmt.Println(testSlice)
	testSlice = append(testSlice, 1)
	testSlice = append(testSlice, 2)
	testSlice = append(testSlice, 3)
	testSlice = append(testSlice, 4)
	testSlice = append(testSlice, 5)
	fmt.Println(testSlice)
	fmt.Println(cap(testSlice))
}
