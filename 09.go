/* 返回参数如果如果命名的话，只需要返回语句即可 */
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}