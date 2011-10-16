/* Go Slices: usage and internals:
http://blog.golang.org/2011/01/go-slices-usage-and-internals.html
*/
package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
