// 在定义map的同时初始化
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 定义一个map对象m，这个map是string对应于Vertex
var m = map[string]Vertex {
	"Bell Labs": Vertex {
		40.68433, 74.39967,
	},
	"Google" : Vertex {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
