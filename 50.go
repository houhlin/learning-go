/* 实际上，可以在自己的包里的任何类型上定义方法，而不仅仅是structs.
你不能在另外一个包里的方法上定义方法，或者在基本类型上定义方法 */
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)	// sqrt2是math包里的常量
	fmt.Println(f.Abs())
}
