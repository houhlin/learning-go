/** range -- 循环访问slice或者map 
第一个是序列，第二个是对应的序列的值
**/


package main

import "fmt"

var  pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
