package safecast_test

import (
	"github.com/chen3feng/safecast"
	"testing"
)

func TestTo(t *testing.T) {
	_, ok := safecast.To[int8](256)
	if ok {
		t.Fail()
	}
}
