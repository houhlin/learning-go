/* slice用make来创建。make分配了初始化为0的数组，返回指向该数组的slice。
slice有length(长度)和capacity(容量)。容量是长度所能达到的最大值，
知道capacity的话，只需要在make里传递第3个参数，如
b := make([]int, 0, 5)		len(b)=0 cap(b)=5
slice也可以增长(取决于capacity)
*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
