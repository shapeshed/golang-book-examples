package example05

import "testing"

func TestStringFromAssignment(t *testing.T) {
	got := StringFromAssignment(10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Fatalf("StringFromAssignment(10) = %q, want %q", got, want)
	}
}

func TestStringFromAppendJoin(t *testing.T) {
	got := StringFromAppendJoin(10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Fatalf("StringFromAppendJoin(10) = %q, want %q", got, want)
	}
}

func TestStringFromBuffer(t *testing.T) {
	got := StringFromBuffer(10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Fatalf("StringFromBuffer(10) = %q, want %q", got, want)
	}
}

func BenchmarkStringFromAssignment(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAssignment(100)
	}
}

func BenchmarkStringFromAppendJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAppendJoin(100)
	}
}

func BenchmarkStringFromBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromBuffer(100)
	}
}
