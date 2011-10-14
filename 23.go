/* Go的基本类型：
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64
complex64 complex128
*/

package main

import (
	"cmath"
	"fmt"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmath.Sqrt(-5+12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
