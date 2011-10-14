// for -- Go语言的唯一遍历结构
// for的结构和C、Java差不多，没有()，强制{}

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
