/**
 * @author: Real
 * Date: 2022/11/18 0:06
 */
package main

import (
	"log"
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
		log.Printf("第 %d 个用户 id 是 %s\n", index, arg)
		uid, err := strconv.Atoi(arg)
		checkErr(err)
		calUserDeliveryFee(uid)
	}
}
