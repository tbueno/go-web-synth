package main

import (
	"strconv"
	"syscall/js"

	"github.com/tbueno/go-web-synth/dom"
	"github.com/tbueno/go-web-synth/synth"
)

func play(i []js.Value) {
	p := dom.Page{}
	audioCtx := p.Get("AudioContext").New()

	s := synth.Synth{Ctx: audioCtx}
	osc := s.CreateOscillator(synth.Sine, 440)
	now, _ := strconv.ParseFloat(audioCtx.Get("currentTime").String(), 32)

	osc.Start(now)
	osc.Stop(now + 0.5)
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
