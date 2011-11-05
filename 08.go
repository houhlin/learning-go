// swap function
package main

import "fmt"

// swap 函数接受2个字符串，返回交换后的字符串
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
