/* if 也可以像for一样在condition前加入语句
需要注意的是: 语句中定义的变量作用范围仅限于if语句本身
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
