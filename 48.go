/* Advanced Exercise: Complex cube roots
For cube roots, Newton's method amounts to repeatings:
z = z - (z^3 - x) / (3*z^2)
Find the cube roots of 2, just to make sure the algorithm works. There is a cmath.Pow function.
update for go 1 */
package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	var z complex128 = 1.0
	for i:=0; i<10; i++ {
		z = z - (cmplx.Pow(z,3)-x)/(3*cmplx.Pow(z,2))
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
