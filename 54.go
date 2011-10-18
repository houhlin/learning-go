/* Errors: An error is anything that can describe itself 
我的理解：任意类型，只要实现String() string就实现了Error接口？
package os
type Error interface {
	String() string		
}
*/

package main

import (
	"fmt"
	"os"
	"time"
)

type MyError struct {
	When *time.Time
	What string
}

// 实现 Error 接口
func (e *MyError) String() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// run函数返回os.Error接口，而这个接口是&MyError(结构体对应的指针)
func run() os.Error {
	return &MyError {
		time.LocalTime(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

