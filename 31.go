// map初始化简写
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 定义一个map对象m，这个map是string对应于Vertex
var m = map[string]Vertex {
	"Bell Labs": { 40.68433, 74.39967, },
	"Google" : { 37.42202, -122.08408, },
}

func main() {
	fmt.Println(m)
}
