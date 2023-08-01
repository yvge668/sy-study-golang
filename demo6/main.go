package main

import "fmt"

func createMap() {
	//使用 make 函数传入键和值的类型，可以创建 map 。具体语法为 make(map[KeyType]ValueType) --> make(map[键类型]值类型)
	// 创建一个键类型为 string 值类型为 int 名为 scores 的 map
	scores := make(map[string]int)
	steps := make(map[string]string)
	fmt.Println(scores)
	fmt.Println(steps)

	//创建并且赋值
	var steps2 map[string]string = map[string]string{
		"呼保义": "宋江",
		"玉麒麟": "卢俊义",
		"豹子头": "林冲",
	}
	fmt.Println(steps2)

	//短定义
	steps3 := map[string]string{
		"浪里白条": "张顺",
		"行者":   "武松",
		"花和尚":  "鲁智深",
	}
	fmt.Println(steps3)
}

func addMapElement() {
	// 可以使用 `map[key] = value` 向 map 添加元素。
	//定义
	scores := make(map[string]int)
	scores["马云"] = 100
	scores["马化腾"] = 199
	scores["玛诗可"] = 103

	fmt.Println(scores)
}

func operationMapElement() {
	//创建并且赋值
	var steps map[string]string = map[string]string{
		"呼保义": "宋江",
		"玉麒麟": "卢俊义",
		"豹子头": "林冲",
	}
	fmt.Println(steps)

	// 若 key 已存在，使用 map[key] = value 可以直接更新对应 key 的 value 值。
	steps["豹子头"] = "潘金莲" //更新
	fmt.Println(steps)

	// 直接使用 map[key] 即可获取对应 key 的 value 值,如果 key不存在,会返回其 value 类型的零值。
	fmt.Println(steps["玉麒麟"]) //通过键获取值

	//使用 delete(map, key)可以删除 map 中的对应 key 键值对,如果 key 不存在,delete 函数会静默处理，不会报错。
	delete(steps, "呼保义") //通过键删除
	fmt.Println(steps)

	//判断key是不是存在
	// 如果我们想知道 map 中的某个 key 是否存在，可以使用下面的语法：value, ok := map[key]
	v3, ok := steps["github"]
	fmt.Println(ok)
	fmt.Println(v3)

	v4, ok := steps["玉麒麟"]
	fmt.Println(ok)
	fmt.Println(v4)
	/*
	   这个语句说明 map 的下标读取可以返回两个值，第一个值为当前 key 的 value 值，第二个值表示对应的 key 是否存在，若存在 ok 为 true ，若不存在，则 ok 为 false
	*/
}

func traverseMap() {
	//遍历map
	var steps map[string]string = map[string]string{
		"呼保义":  "宋江",
		"玉麒麟":  "卢俊义",
		"豹子头":  "林冲",
		"浪里白条": "张顺",
		"行者":   "武松",
		"花和尚":  "鲁智深",
	}

	for key, value := range steps {
		fmt.Println(key + "--->" + value)
	}

	fmt.Println("steps这个map中有", len(steps), "个键值对") // 4

}

func mapByReference() {
	//当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量。

	steps4 := map[string]string{
		"呼保义":  "宋江",
		"玉麒麟":  "卢俊义",
		"豹子头":  "林冲",
		"浪里白条": "张顺",
		"行者":   "武松",
		"花和尚":  "鲁智深",
	}
	fmt.Println("steps4: ", steps4)

	newSteps4 := steps4
	newSteps4["呼保义"] = "996ICU"
	newSteps4["玉麒麟"] = "github搬运工"
	newSteps4["行者"] = "辛辛苦苦打工人"
	fmt.Println("steps4: ", steps4)

	fmt.Println("newSteps488: ", newSteps4)

}

func main() {
	createMap()
	addMapElement()
	operationMapElement()
	traverseMap()
	mapByReference()
}
