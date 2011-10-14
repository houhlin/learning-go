// go有指针，但是没有指针运算
// 结构体成员可以通过指针来访问，通过指针间接访问是透明的
package main
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}
