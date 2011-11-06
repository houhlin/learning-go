/* Exercise: Fibonacci closure
Implement a fibonacci function that returns a function(a closure) that return successive fibonacci numbers.
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946，………………
*/

package main

import "fmt"

// fibonacci is a function that returns a function that returns a int
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// outputs: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
