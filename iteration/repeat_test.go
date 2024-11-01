package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("Repeated result was incorrect, got: %s, want: %s.", repeated, expected)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 7)
	fmt.Println(repeated)

	// Output: aaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
