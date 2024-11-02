package di

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("MyGreeterHandler called with method %s \n", r.Method)
	Greet(w, "world!")
}
