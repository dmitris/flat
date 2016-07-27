package main

import (
	"github.com/orloffm/cat"
)

func main() {
	c := new(cat.Cat)
	t := new(toy.RubberToy)

	c.GentlyPushAToy(*t)
}
