package main

import "fmt"

func main() {
	m1 := make(map[string]string) // map[键类型] 值类型

	m1["name"] = "Jason" // 键值对赋值
	m1["age"] = "23"
	m1["sex"] = "male"
	fmt.Println(m1)

	//按键取值
	fmt.Println(m1["name"])         // name
	fmt.Println(m1["gender"])       // 空字符串
	fmt.Println(m1["gender"] == "") // true

	//删除某值
	delete(m1, "age") // 如果不存在age键，则也不会报错
	// 如果需要清空映射，直接分配新的内存就行
	m1 = make(map[string]string)
	fmt.Println("after clear , map =", m1)

	// 三种定义 map 的方式
	// 1、先声明再make
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	// 2、直接make
	map2 := make(map[string]string, 10)
	// 3、声明时赋值
	map3 := map[string]string{
		"1": "first",
		"2": "second", // 逗号不能省略，否则编译错误
	}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	// 遍历 Map
	for k, v := range map3 {
		fmt.Println("Key:", k, "Value:", v)
	}

	// 切片动态改变 map 的个数
	var sliceMap []map[string]string
	sliceMap = make([]map[string]string, 10)
	sliceMap = append(sliceMap, map2, map3)
	fmt.Println("sliceMap: ", sliceMap)

	// 判断 key 是否存在
	var dict map[string]string
	dict = make(map[string]string, 5)
	dict["test1"] = "测试一"
	dict["test2"] = "测试二"
	dict["test3"] = "测试三"
	key := "test1"
	value, contains := dict[key]
	if contains {
		fmt.Println("存在：", value)
	} else {
		fmt.Println("不存在 Key: ", key)
	}

}
