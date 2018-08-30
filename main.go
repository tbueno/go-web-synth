package main

import (
	"strconv"
	"syscall/js"
)

type Synth struct {
	ctx js.Value
}

func (s Synth) createOscillator(t string, freq uint) js.Value {
	osc := s.ctx.Call("createOscillator")
	osc.Set("type", t)
	osc.Get("frequency").Set("value", freq)
	osc.Call("connect", s.ctx.Get("destination"))
	return osc
}

func play(i []js.Value) {
	audioCtx := js.Global().Get("AudioContext").New()
	s := Synth{ctx: audioCtx}
	osc := s.createOscillator("sine", 440)
	now, _ := strconv.ParseFloat(audioCtx.Get("currentTime").String(), 32)

	osc.Call("start", now)
	osc.Call("stop", now+0.5)
}

func registerCallbacks() {
	js.Global().Set("play", js.NewCallback(play))
}

func main() {
	// Guarantees that the program does not exit before it is ready
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
