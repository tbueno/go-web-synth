package synth

import "syscall/js"

type Oscillator struct {
	nodeOsc js.Value
}

type WaveType string

const (
	Sine     = "sine"
	Square   = "square"
	Triangle = "triangle"
	Sawtooth = "sawtooth"
)

func (o Oscillator) Start(t float64) {
	o.nodeOsc.Call("start", t)
}

func (o Oscillator) Stop(t float64) {
	o.nodeOsc.Call("stop", t)
}

func NewOscillator(s Synth, t string, f uint) Oscillator {
	osc := s.Ctx.Call("createOscillator")
	osc.Set("type", t)
	osc.Get("frequency").Set("value", f)
	osc.Call("connect", s.Ctx.Get("destination"))

	return Oscillator{nodeOsc: osc}
}
