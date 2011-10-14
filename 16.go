// 像C和Java那样,可以把初始化和结束语句留空
// 这样难道每次自动+1?
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000;  {
		sum += sum
	}
	fmt.Println(sum)
}
