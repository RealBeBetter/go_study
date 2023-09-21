package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 数组的定义
	var array [5]float64
	// 总数
	sum := 0.0

	rand.Seed(time.Now().Unix())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Float64()*20 + 5
		fmt.Println("数组：第", i+1, "个数是：", array[i])
		sum += array[i]
	}

	fmt.Println("总和为：", sum)
	// 浮点数运算，需要将整数转换为浮点数
	fmt.Println("均值为：", sum/float64(len(array)))

	// 数组的初始化
	var nums1 [4]int = [4]int{1, 2, 3, 4}
	var nums2 = [4]int{1, 2, 3, 4}
	// 自行判断长度，中括号里...一个不能少
	var nums3 = [...]int{1, 2, 3, 4}
	// 指定索引和值, 格式 index: number
	var nums4 = [...]int{1: 3, 0: 4, 2: 5, 5: 8}
	fmt.Println("数组 nums1 的值为：", nums1)
	fmt.Println("数组 nums2 的值为：", nums2)
	fmt.Println("数组 nums3 的值为：", nums3)
	fmt.Println("数组 nums4 的值为：", nums4)

	// 传递数组作为指针，修改数组的值
	modify(&array)
	// 根据输出结果可以得知，指针操作属于值传递
	fmt.Println(array)

	sum = 0.0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	fmt.Println("均值为：", sum/(float64(len(array))))

}

func modify(array *[5]float64) {
	array[0] += 5
}
