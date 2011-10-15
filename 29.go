/* map -- key to value, map[key]value
初始化map的时候使用make而不是new
取得对应key的value值使用map[key]
*/
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex	// 定义一个map对象m，这个map是string对应于Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.68433, 74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
