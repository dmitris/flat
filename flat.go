package flat

import (
	"github.com/orloffm/cat"
	"github.com/orloffm/dog"
	"github.com/orloffm/toy"
)

func main() {
	c := new(cat.Cat)
	d := new(dog.Dog)

	t := new(toy.RubberToy)

	c.GentlyPush(*t)
	d.Squeeze(*t)
}
