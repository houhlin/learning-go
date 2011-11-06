/* Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsighed integers. When you run theprogram, it will display your picture, interpreting the integers as grayscale(well, bluescale) values.
The choice of images is up to you. Interesting functions inlcude x^y, (x+y)/2, and x*y.
(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
*/

package main
import (
//	"go-tour.googlecode.com/hg/pic"
	"tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var r = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		r [i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			r[i][j] = uint8(i*j)
		}
	}
	return r
}

func main() {
	pic.Show(Pic)
}
