/* 
在结构体赋值中可以使用Name:syntax的形式来对结构体的某一成员单独赋值，结构体成员顺序不限
*/
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // 隐含 Y:0 
	s = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
