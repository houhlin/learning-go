/* A type implements an interface by implementing the methods.*/

package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err os.Error)
}

type Writer interface {
	Write(b []byte) (n int, err os.Error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

