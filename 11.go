/** 
var 初始化的同时可以赋值，如果初始化存在的话，可以省略类型名，变量会根据初始化的数据自动决定类型
**/
package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, java, python = true, false, "No"

func main() {
	fmt.Println(x, y, z, c, java, python)
}
