package foo

import "testing"

func TestTruth(t *testing.T) {
	if false != true {
		t.Fatal("The world is crumbling")
	}
}
