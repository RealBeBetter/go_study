package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// StuStruct 1、声明结构体
type StuStruct struct {
	Name  string
	Sex   string
	Score int
}

// StuSlice 2、声明一个结构体切片
type StuSlice []StuStruct

// Len 3、实现Len()、Less()、Swap()三个接口
func (stu StuSlice) Len() int {
	return len(stu)
}

func (stu StuSlice) Less(i, j int) bool {
	// 按照学生成绩排序 从大到小排序，从小到大用 <
	return stu[i].Score > stu[j].Score
}

func (stu StuSlice) Swap(i, j int) {
	//temp := stu[i]
	//stu[i] = stu[j]
	//stu[j] =temp
	// 上面三行等价于
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	var stuSlice StuSlice

	for i := 0; i < 5; i++ {
		stu := StuStruct{fmt.Sprintf("学生%d", i), "男", rand.Intn(100)}
		stuSlice = append(stuSlice, stu)
	}
	fmt.Printf("排序前数据：%v \n", stuSlice)
	sort.Sort(stuSlice)
	fmt.Printf("排序后数据：%v \n", stuSlice)

}
