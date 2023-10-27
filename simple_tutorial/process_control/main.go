/**
 * @author wei.song
 * @since 2023/10/27 14:57
 */
package main

import "fmt"

func main() {
	// process control
	fmt.Println("---------process control condition--------")
	processControlCondition()
	fmt.Println("---------process control switch--------")
	processControlSwitch()
	fmt.Println("---------process control for--------")
	processControlFor()
}

func processControlCondition() {
	age := 18
	if age < 18 {
		fmt.Println("Kid")
	} else {
		fmt.Println("Adult")
	}

	// 可以简写为
	if ageValue := 18; ageValue < 18 {
		fmt.Println("Kid")
	} else {
		fmt.Println("Adult")
	}
}

type Gender int

const (
	Male Gender = iota
	Female
)

func processControlSwitch() {
	// male
	switch gender := Male; gender {
	case 0:
		fmt.Println("male")
	case 1:
		fmt.Println("female")
	default:
		fmt.Println("unknown")
	}

	// female
	// male
	// unknown
	switch gender := Female; gender {
	case Female:
		fmt.Println("female")
		fallthrough
	case Male:
		fmt.Println("male")
		fallthrough
	default:
		fmt.Println("unknown")
	}
	// 使用 fallthrough 表示会继续执行下一个 case

}

func processControlFor() {
	// general for-loop
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}

	// array for-loop
	nums := []int{10, 20, 30, 40, 50}
	for index, num := range nums {
		fmt.Println("current index, number: ", index, num)
	}

	mapFor := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	for key, value := range mapFor {
		fmt.Println("current key, value: ", key, value)
	}

}
