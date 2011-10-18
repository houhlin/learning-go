/* Web servers: http包，实现http.handler
package http

type Handler interface {
	ServerHTTP( w ResponseWriter,
	r *Request)
}
*/

package main

import (
	"fmt"
	"http"
)

type Hello struct {}

func (h Hello) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, "Hello!")
	}

	func main() {
		var h Hello
		http.ListenAndServe("localhost:4000",h)
	}
