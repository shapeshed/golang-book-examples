// Package example03 shows how to use the godoc tool.
package example03

import (
	"errors"
)

// Animal specifies an animal
type Animal struct {
	Name string // Name holds the name of an Animal.

	// Age holds the name of an Animal.
	Age int
}

// ErrNotAnAnimal is returned if the name field of the Animal struct is Human.
var ErrNotAnAnimal = errors.New("Name is not an animal")

// Hello sends a greeting to the animal.
func (a Animal) Hello() (string, error) {
	if a.Name == "Human" {
		return "", ErrNotAnAnimal
	}
	s := "Hello " + a.Name
	return s, nil
}
