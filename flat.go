package main

import (
	"github.com/orloffm/cat"
	"github.com/orloffm/toy"
)

func main() {
	c := new(cat.Cat)
	t := new(toy.RubberToy)

	c.GentlyPushAToy(*t)
}
