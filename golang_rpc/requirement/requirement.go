/**
 * @author Real
 * @since 2023/11/21 21:23
 */
package main

import "log"

type Result struct {
	Num, Ans int
}

type Calc int

func (cal *Calc) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	calc := new(Calc)
	var result Result
	_ = calc.Square(11, &result)
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
