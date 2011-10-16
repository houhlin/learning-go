// slice -- 切片,从0开始
// s[lo:hi] 区间范围是[lo, high-1]
package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])
	
	// 缺失low index默认是0
	fmt.Println("p[:3] ==", p[:3])
	
	// 缺失high index默认是len(s)
	fmt.Println("p[4:] ==", p[4:])
}