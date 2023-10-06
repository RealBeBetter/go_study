/**
 * @author: Real
 * Date: 2022/11/18 0:06
 */
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	execSql()
	args := os.Args
	for index, arg := range args {
		if index == 0 {
			continue
		}
		fmt.Println("第", index, "个用户 id 是 ", arg)
		uid, err := strconv.Atoi(arg)
		checkErr(err)
		calUserDeliveryFee(uid)
	}
}
