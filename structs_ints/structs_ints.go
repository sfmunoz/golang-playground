//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Tue Oct  3 05:11:38 PM UTC 2023
//

// {{{ package

package structs_ints

// }}}
// {{{ imports

import "fmt"
import "math"

// }}}
// ---- types: interfaces ----
// {{{ type Shape interface

type Shape interface {
	area() float64
}

// }}}
// ---- types: structs ----
// {{{ type Rect struct

type Rect struct {
	h float64
	w float64
}

// }}}
// {{{ type Circ struct

type Circ struct {
	r float64
}

// }}}
// ---- funcs (private) ----
// {{{ func (r Rect) area()

func (r Rect) area() float64 {
	return r.h * r.w
}

// }}}
// {{{ func (c Circ) area()

func (c Circ) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

// }}}
// {{{ func getArea(s Shape)

func getArea(s Shape) float64 {
	return s.area()
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	r := Rect{4, 5}
	c := Circ{4}
	fmt.Printf("Rect%+v -> %.2f -- %.2f\n", r, r.area(), getArea(r))
	fmt.Printf("Circ%+v -> %.2f -- %.2f\n", c, c.area(), getArea(c))
}

// }}}
