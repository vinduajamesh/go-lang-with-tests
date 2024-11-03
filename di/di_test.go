package di

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	fmt.Printf("cheese %v", buffer)
	Greet(&buffer, "James")
	got := buffer.String()
	want := "Hello, James"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
