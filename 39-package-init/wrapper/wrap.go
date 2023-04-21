package wrapper

import (
	"demo/shape"
	"fmt"
)

func New(iface shape.IShape) *Wrapper {
	w := new(Wrapper)
	w.Iface = iface
	return w
}

type Wrapper struct {
	Iface shape.IShape
}

func (w *Wrapper) Area() {
	fmt.Println(w.Iface.Area())
}

func (w *Wrapper) Perimeter() {
	fmt.Println(w.Iface.Perimeter())
}
