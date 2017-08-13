package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func main() {

	r := R2D2{
		Broken: true,
	}

	err := r.PowerOn()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

}
