//function 函数, 类型在变量名后面 
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42,13))
}
