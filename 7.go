/* 注意：这里 x int, y int 可以简写为 x, y int */
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42,13))
}
