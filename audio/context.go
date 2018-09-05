package audio

import (
	"github.com/tbueno/go-web-synth/dom"
)

type Context struct {
	Ctx dom.Node
}

// CreateOscillator builds an oscillator and connect it to the synth context
func (c Context) CreateOscillator(t string, freq uint) Oscillator {
	return NewOscillator(c, t, freq)
}
