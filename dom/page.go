package dom

import (
	"syscall/js"
)

type Node struct {
	elem js.Value
}

func NewNode(elem js.Value) Node {
	return Node{elem}
}

func (n Node) Value() js.Value {
	return n.elem
}

func (n Node) Get(content string) Node {
	return Node{n.elem.Get(content)}
}

func (n Node) Set(attr string, v interface{}) {
	n.elem.Set(attr, v)
}

func (n Node) New() Node {
	return Node{n.elem.New()}
}

func (n Node) Call(f string, args ...interface{}) js.Value {
	return n.elem.Call(f, args)
}

func (n Node) String() string {
	return n.elem.String()
}

type Page struct {
}

func (p Page) Get(content string) Node {
	return Node{js.Global().Get(content)}
}

func (p Page) Call(f string) Node {
	return Node{js.Global().Call(f)}
}
