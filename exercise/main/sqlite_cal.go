/**
 * @author: Real
 * Date: 2022/11/17 20:15
 */
package main

import (
	"fmt"
	"math"
)

const startFee float64 = 18.0
const startWeight float64 = 1.0
const insuranceRate float64 = 0.01

func main() {
	weight := 211.0
	fee := CalDeliveryFee(weight)
	fmt.Println(fee)
}

// CalDeliveryFee /** 1.快递费用按重量W计算，其中包含运费和保险费;
//2.1KG(含)以内18元，包含运费和保险费;
//3.每增加1kg，运费增加5元，保险费为上一档费用 (运费+保险费)的1%;
//4，每单快递费用计算结果按照四舍五入取整数值
func CalDeliveryFee(weight float64) float64 {
	fmt.Println("begin calculate , weight is", weight)
	if weight <= startWeight {
		return startFee
	} else {
		fee := CalDeliveryFee(weight - 1)
		return math.Round(fee + fee*insuranceRate + 5)
	}
}

// Fee 重量、运费计费
type Fee struct {
	Express   float64
	Insurance float64
}
