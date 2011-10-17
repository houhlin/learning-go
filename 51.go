/* 方法可以关联一个命名的类型或者一个命名类型的指针.
使用指针接受者的2个原因：
1. 避免每次方法调用的适合复制值(当大数组的适合很有效率)
2. 用指针可以让方法可以修改接受者指向的值
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (p *Vertex) Scale(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

func (p *Vertex) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := &Vertex{3, 4}
	p.Scale(5)
	fmt.Println(p, p.Abs())
}
