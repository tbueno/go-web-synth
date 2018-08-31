package synth

import "github.com/tbueno/go-web-synth/dom"

type Oscillator struct {
	nodeOsc dom.Node
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

// func (o Oscillator) Play(d float64) {
// 	now, _ := strconv.ParseFloat(audioCtx.Get("currentTime").String(), 32)
// 	o.nodeOsc.Call("start", t)
// }

func NewOscillator(s Synth, t string, f uint) Oscillator {
	osc := s.Ctx.Call("createOscillator")
	osc.Set("type", t)
	osc.Get("frequency").Set("value", f)
	osc.Call("connect", s.Ctx.Get("destination").Value())

	return Oscillator{nodeOsc: dom.NewNode(osc)}
}
