package example03

import "testing"

func TestTranslation(t *testing.T) {
	got := translate("fr")
	want := "Bonjour "
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestGreeting(t *testing.T) {
	got := Greeting("George", "en-US")
	want := "Hello George"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
