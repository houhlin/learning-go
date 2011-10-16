/** switch 没有条件，代表true **/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.LocalTime()
	switch {
	case t.Hour < 12:
		fmt.Println("Good morning")
	case t.Hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
