package synth

import (
	"github.com/tbueno/go-web-synth/dom"
)

// Synth represents a simple synthetizer connected to an audio context
type Synth struct {
	Ctx dom.Node
}

// CreateOscillator builds an oscillator and connect it to the synth context
func (s Synth) CreateOscillator(t string, freq uint) Oscillator {
	return NewOscillator(s, t, freq)
}
