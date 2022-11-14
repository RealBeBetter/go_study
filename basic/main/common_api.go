/**
 * @author: Real
 * Date: 2022/11/15 1:37
 */
package main

import (
	"fmt"
	"strconv"
	"time"
)

var goTime = "2022-11-15 01:37:05"

func main() {
	//字符串转 time
	str := "2022-12-01 15:00:00"

	res1, _ := time.ParseInLocation(goTime, str, time.Local)
	res2, _ := time.Parse(goTime, str)

	fmt.Printf("res1 = %v\n", res1)
	fmt.Printf("res2 = %v\n", res2)

	//time转字符串
	format1 := res1.Format(goTime)
	format2 := res2.Format(goTime)

	fmt.Printf("format1 = %v\n", format1)
	fmt.Printf("format2 = %v\n", format2)

	//将常规的数据类型转换为string
	str = fmt.Sprintf("%d%c", 123, `a`)

	// int转字符串
	intSize := strconv.Itoa(123)
	fmt.Printf("%T  \n", intSize)
	atoi, _ := strconv.Atoi(intSize)
	// 字符串转int
	fmt.Printf("%T  \n", atoi)
	// 将int64转字符串
	formatInt := strconv.FormatInt(int64(123), 10)
	fmt.Printf("%T  \n", formatInt)
	// 将字符串转换成 10进制64位int 即int64
	parseInt, _ := strconv.ParseInt("123", 10, 64)
	fmt.Printf("%T  \n", parseInt)
}

func StrToTime(str string) *time.Time {
	t, _ := time.ParseInLocation("2022-11-15 01:12:05", str, time.Local)
	return &t
}

func TimeToStr(t *time.Time) string {
	return t.Format("2022-11-15 01:23:05")
}
