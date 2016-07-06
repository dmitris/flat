package main

import (
	"github.com/orloffm/cat"
	"github.com/orloffm/dog"
)

func main() {
	c := new(cat.Cat)
	d := new(dog.Dog)
	
	t := d.BringAToy()

	c.GentlyPushAToy(*t)
}
