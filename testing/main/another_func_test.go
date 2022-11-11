/**
 * @author: Real
 * Date: 2022/11/12 0:56
 */
package main

import (
	"fmt"
	"testing"
)

func TestAddSale(t *testing.T) {
	// 测试时调用其他的方法
	sale := &Sale{WidgetId: 9, Qty: 80, Street: "Huanghe South Road", City: "Anyang Henan", State: "China", Zip: 455000, SaleDate: "2020-03-24"}
	sale.AddSale()
	// 执行 Run 方法，可以在测试时执行另一个测试函数
	t.Run("fun", testSale)

	// 所有运行结果
	//C:\Users\Real\GoProjects\go_study\testing\main>go test -v
	//Testing start...
	//=== RUN   TestAdd
	//    add_test.go:17: calculate succeed, your answer is 55
	//--- PASS: TestAdd (0.00s)
	//=== RUN   TestAddSale
	//=== RUN   TestAddSale/fun
	//test function
	//--- PASS: TestAddSale (0.00s)
	//    --- PASS: TestAddSale/fun (0.00s)
	//=== RUN   TestSub
	//    sub_test.go:22: sub calculate error, correct answer is 0
	//--- FAIL: TestSub (0.00s)
	//FAIL
	//exit status 1
	//FAIL    go-study/testing/main   0.473s
}

type Sale struct {
	WidgetId int
	Qty      int
	Street   string
	City     string
	State    string
	Zip      int
	SaleDate string
}

func (sale *Sale) AddSale() int {
	return sale.Qty
}
func testSale(t *testing.T) {
	fmt.Println("test function")
}
