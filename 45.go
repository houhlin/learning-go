/* exercise: Maps 
Implement WordCount. It should return a map of the counts of each "word" in the string s. The wc.Test func runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful
*/

package main
import (
//	"go-tour.googlecode.com/hg/wc"	
	"tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
//	var out []string			// slices
	out := strings.Fields(s)
	for _, v := range out {
		if _, ok := m[v]; ok {			//  测试m[v]是否在out[]里, ok为true时
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
