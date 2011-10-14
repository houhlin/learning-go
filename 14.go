/* 数字常量是高精度数。未定义类型的常量有上下文决定
这意味着自动转换???
*/
package main

import "fmt"

const (
	Big = 1<<100
	Small = Big>>99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x*0.1
}

func main() {
	fmt.Println(needInt(Small))
	// 溢出 overflow
//	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

