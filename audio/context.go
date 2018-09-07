package audio

import (
	"syscall/js"

	"github.com/tbueno/go-web-synth/dom"
)

type Context struct {
	Ctx dom.Node
}

type AudioDestinationNode struct {
	MaxChannelCount uint
	node            dom.Node
}

func (adn AudioDestinationNode) Connect(dest Node) {
	// The destination should not connect to anything.
}

func (adn AudioDestinationNode) Value() js.Value {
	return adn.node.Value()
}

// CreateOscillator builds an oscillator and connect it to the synth context
func (c Context) CreateOscillator() Oscillator {
	return NewOscillator(c)
}

func (c *Context) Destination() Node {
	return AudioDestinationNode{node: c.Ctx.Get("destination")}
}
