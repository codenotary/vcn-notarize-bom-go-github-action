package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 1
	figure.NewColorFigure(babbler.Babble(), "", "green", true).Print()
}
