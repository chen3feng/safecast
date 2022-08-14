package safecast_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/chen3feng/safecast"
)

func TestTo(t *testing.T) {
	_, ok := safecast.To[int8](256)
	if ok {
		t.Fail()
	}
}

func ExampleTo_intNoOverflow() {
	b, ok := safecast.To[byte](255)
	fmt.Print(b, ok)
	// Output:
	// 255 true
}

func ExampleTo_intOverflow() {
	b, ok := safecast.To[byte](256)
	fmt.Print(b, ok)
	// Output:
	// 0 false
}

func ExampleTo_valueInRange() {
	n, ok := safecast.To[uint](1)
	fmt.Print(n, ok)
	// Output:
	// 1 true
}

func ExampleTo_valueOutOfRange() {
	n, ok := safecast.To[uint32](-1)
	fmt.Print(n, ok)
	// Output:
	// 4294967295 false
}

func ExampleTo_floatOverflow() {
	n, ok := safecast.To[float32](math.MaxFloat32 * 2)
	fmt.Print(n, ok)
	// Output:
	// +Inf false
}
