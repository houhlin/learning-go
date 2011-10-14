/** 在函数内部， := 赋值表达式;在函数外部，:= 不可用 **/

package main

import "fmt"

func main() {
	var x, y, z int = 1, 2, 3
	c, java, python := true, false, "No"
	fmt.Println(x, y, z, c, java, python)
}
