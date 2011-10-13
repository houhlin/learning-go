/* 
每个go程序都是有package组成的
包的名字是导入路径的最后一个元素
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Happy", math.Pi, "day")
}
