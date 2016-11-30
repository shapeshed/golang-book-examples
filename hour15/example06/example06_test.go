package salt

import "testing"

func TestSaltSecret(t *testing.T) {
	got := SaltSecret("foo")
	want := "fooiep4Og#aiW3f\n"
	if got != want {
		t.Fatalf("SaltSecret(\"foo\") = %q, want %q", got, want)
	}
}

func BenchmarkSaltSecret(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SaltSecret("somesecret")
	}
}
