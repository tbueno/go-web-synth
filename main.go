package main

import (
	"syscall/js"

	"github.com/tbueno/go-web-synth/audio"
	"github.com/tbueno/go-web-synth/dom"
)

func play(i []js.Value) {
	p := dom.NewPage()

	audioCtx := audio.Context{Ctx: p.Get("AudioContext").New()}
	osc := audioCtx.CreateOscillator()

	osc.Connect(audioCtx.Destination())

	osc.Type(audio.Sawtooth)
	osc.Frequency(440)

	osc.Play(0.2)

	counter := p.FindByID("counter")
	counter.InnerHTML("0")
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
