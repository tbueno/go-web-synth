package audio

import (
	"strconv"
	"syscall/js"

	"github.com/tbueno/go-web-synth/dom"
)

// Node represents an audioNode from Web Audio API.
type Node interface {
	Connect(dest Node)
	Value() js.Value
}

type Oscillator struct {
	nodeOsc dom.Node
	osc     js.Value
}

type WaveType string

const (
	Sine     = "sine"
	Square   = "square"
	Triangle = "triangle"
	Sawtooth = "sawtooth"
)

func NewOscillator(c Context) Oscillator {
	osc := c.Ctx.Call("createOscillator")
	return Oscillator{nodeOsc: dom.NewNode(osc), osc: osc}
}

func (o *Oscillator) Type(t string) {
	o.nodeOsc.Set("type", t)
}

func (o *Oscillator) Frequency(f uint) {
	o.nodeOsc.Set("frequency", f)
}

func (o *Oscillator) Play(d float64) {
	n, _ := strconv.ParseFloat(o.nodeOsc.Get("currentTime").String(), 32)
	o.nodeOsc.Call("start", n)
	o.nodeOsc.Call("stop", d)
}

func (o Oscillator) Start(t float64) {
	o.nodeOsc.Call("start", t)
}

func (o Oscillator) Stop(t float64) {
	o.nodeOsc.Call("stop", t)
}

func (o *Oscillator) Connect(dest Node) {
	o.osc.Call("connect", dest.Value())
}
