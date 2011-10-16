/* switch语句 -- case语句自动breka，除非以fallthrough结尾 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.", os)
	}
}
