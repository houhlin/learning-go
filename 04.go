/* fmt.Printf -- 带format string */
package main

import "fmt"
import "math"

func main() {
	fmt.Printf("Now you have have %g Problems.", math.Nextafter(2, 3))
}
