package example01

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Fatal("The world is crumbling")
	}
}
