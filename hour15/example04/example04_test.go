package example04

import "testing"

type GreetingTest struct {
	name   string
	locale string
	want   string
}

var greetingTests = []GreetingTest{
	{"George", "en-US", "Hello George"},
	{"Chloé", "fr-FR", "Bonjour Chloé"},
	{"Giuseppe", "it-IT", "Ciao Giuseppe"},
}

func TestGreeting(t *testing.T) {
	for _, test := range greetingTests {
		got := Greeting(test.name, test.locale)
		if got != test.want {
			t.Errorf("Greeting(%s,%s) = %v; want %v", test.name, test.locale, got, test.want)
		}
	}
}
