/* 练习: 用Newton's method实现开根号 
Newton's method: 选择一个开始点z，重复z=z-(z^2-x)/2z, 10次后测试结果，或者更多次后查看结果，与math.Sqrt相比较
*/

package main

import (
	"fmt"
	"math"
)

// 实现这个函数，关键是数学思想
func Sqrt(x float64) float64 {
	var z = 1.0
	for i := 1; i < 10; i++ {		// i<6就很接近了
		z = z - 0.5*(z*z -x)/z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
