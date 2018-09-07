package dom

import (
	"syscall/js"
)

type Page struct {
	doc js.Value
}

func NewPage() Page {
	return Page{doc: js.Global().Get("document")}
}

func (p Page) FindByID(id string) Node {
	return Node{p.doc.Call("getElementById", id)}
}

func (p Page) Get(content string) Node {
	return Node{js.Global().Get(content)}
}

func (p Page) Call(f string) Node {
	return Node{js.Global().Call(f)}
}
