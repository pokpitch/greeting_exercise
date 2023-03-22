package greeting

import "fmt"

func greeting(name string) string {
	return "Hello, " + name
}

func greetingWithAge(name string, age int) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

func greetingWithAgeAndDrink(name string, age int, drink string) string {
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s", name, age, drink)
}
