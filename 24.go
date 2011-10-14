// 结构体 struct, type 关键字
// 结构体用{}包围，打印的适合go很smart
package main
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
