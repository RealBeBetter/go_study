/**
 * @author: Real
 * Date: 2022/11/17 22:46
 */
package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestExec(t *testing.T) {
	// 编写单元测试
	Exec()
	t.Logf("calculate succeed...")
}

func Exec() {
	execSql()
	args := os.Args
	for index, arg := range args {
		fmt.Sprintf("第 %d 个用户 id 是 %s", index+1, arg)
		uid, err := strconv.Atoi(arg)
		checkErr(err)
		calUserDeliveryFee(uid)
	}
}
