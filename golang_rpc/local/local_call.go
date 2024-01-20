/**
 * @author Real
 * @since 2023/11/20 23:30
 */
package main

import (
	"golang.org/x/tools/container/intsets"
	"log"
)

type Result struct {
	Num, Ans int
}

type Calc int

// CalcSquare 计算面积
//  @receiver cal
//  @param num
//  @return *Result
func (cal *Calc) CalcSquare(num int) *Result {
	if num > intsets.MaxInt/2 {
		panic("error argument: num is too big......")
	}

	return &Result{
		Num: num,
		Ans: num * num,
	}
}

func main() {
	calc := new(Calc)
	square := calc.CalcSquare(10)
	log.Printf("%d^2 = %d", square.Num, square.Ans)
}
