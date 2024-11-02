package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vinduajamesh/go-lang-with-tests/di"
)

func main() {
	fmt.Println("Hello World")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
}
