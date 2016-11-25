package pariszoo

import (
	"fmt"
	"io"

	"github.com/sbinet/envol-2016-go/zoo"
)

var gAnimals []zoo.Animal

func Add(animals []zoo.Animal) {
	for _, animal := range animals {
		gAnimals = append(gAnimals, animal)
	}
}

type Fish struct{}

func (f *Fish) Speak(w io.Writer) error {
	fmt.Fprintf(w, "mute, mute\n")
	return nil
}

func (f *Fish) NumAppendage() int {
	return 0 // heh...
}

// explicitly tell *Fish implements zoo.Animal
var _ zoo.Animal = (*Fish)(nil)
