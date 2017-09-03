package example06

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("George")
	want := "Hello George"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
