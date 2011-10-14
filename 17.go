// 这个更猛，for可以像while一样用(golang里面没有while)
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

