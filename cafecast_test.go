package safecast_test

import (
	"github.com/chen3feng/safecast"
	"testing"
)

func TestXxxToXxx(t *testing.T) {
	n, ok := safecast.Int16ToUint(11)
	if !ok || n != 11 {
		t.Fail()
	}
}

func TestTo(t *testing.T) {
	_, ok := safecast.To[int8](256)
	if ok {
		t.Fail()
	}
}
