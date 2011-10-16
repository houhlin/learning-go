/* switch 从上之下运行，当case成功的时候停止 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	// time类的LocalTime函数返回time指针，type time有Weekday的filed
	today := time.LocalTime().Weekday
	switch time.Saturday {
	case today+0:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tomorrow.")
	case today+2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

