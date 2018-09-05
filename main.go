package main

import (
	"strconv"
	"syscall/js"

	"github.com/tbueno/go-web-synth/audio"
	"github.com/tbueno/go-web-synth/dom"
)

func play(i []js.Value) {
	p := dom.NewPage()
	audioCtx := p.Get("AudioContext").New()

	s := audio.Context{Ctx: audioCtx}
	osc := s.CreateOscillator(audio.Sine, 440)
	now, _ := strconv.ParseFloat(audioCtx.Get("currentTime").String(), 32)
	osc.Start(now)
	osc.Stop(now + 0.2)

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
