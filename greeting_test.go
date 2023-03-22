package greeting

import "testing"

func TestGreeting(t *testing.T) {
	given := "Pokpitch"
	want := "Hello, Pokpitch"

	result := greeting(given)
	if result != want {
		t.Errorf("greeting(%q) = %q, want %q", given, result, want)
	}
}

func TestGreetingWithAge(t *testing.T) {
	givenName := "Pokpitch"
	givenAge := 39
	want := "Hello, Pokpitch. You are 39 years old."

	result := greetingWithAge(givenName, givenAge)
	if result != want {
		t.Errorf("greetingWithAge(%q, %d) = %q, want %q", givenName, givenAge, result, want)
	}
}

func TestGreetingWithAgeAndDrink(t *testing.T) {
	givenName := "Pokpitch"
	givenAge := 39
	givenDrink := "Coke"
	want := "Hello, Pokpitch. You are 39 years old and your favorite drink is Coke"

	result := greetingWithAgeAndDrink(givenName, givenAge, givenDrink)
	if result != want {
		t.Errorf("greetingWithAgeAndDrink(%q, %d, %q) = %q, want %q", givenName, givenAge, givenDrink, result, want)
	}
}
