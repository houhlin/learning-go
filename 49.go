/*
	Go没有class，但是可以在结构体类型上面定义方法
	方法的接受方出现在func关键词和函数名之间，如下定义vertex类型上的Abs方法(大写代表可以导出)
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (p *Vertex) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := &Vertex{3, 4}
	fmt.Println(p.Abs())
}