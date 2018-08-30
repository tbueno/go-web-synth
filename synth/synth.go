package synth

import "syscall/js"

// Synth represents a simple synthetizer connected to an audio context
type Synth struct {
	Ctx js.Value
}

// CreateOscillator builds an oscillator and connect it to the synth context
func (s Synth) CreateOscillator(t string, freq uint) Oscillator {
	return NewOscillator(s, t, freq)
}
