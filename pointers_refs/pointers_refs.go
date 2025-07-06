//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Tue Oct  3 04:46:23 PM UTC 2023
//

// {{{ package

package pointers_refs

// }}}
// {{{ imports

import "fmt"

// }}}
// ---- types ----
// {{{ type Data struct

type Data struct {
	val    int
	prints int
}

// }}}
// ---- funcs (private) ----
// {{{ func valChange()

func valChange(x int) {
	prefix := "valChange():"
	fmt.Println(prefix, "x =", x, "/ &x =", &x, "(before)")
	x = 2
	fmt.Println(prefix, "x =", x, "/ &x =", &x, "(after)")
}

// }}}
// {{{ func refChange()

func refChange(x *int) {
	prefix := "refChange():"
	fmt.Println(prefix, "*x =", *x, "/ x =", x, "(before)")
	*x = 2
	fmt.Println(prefix, "*x =", *x, "/ x =", x, "(after)")
}

// }}}
// {{{ func(d Data) valChange()

func (d Data) valChange(val int) {
	d.val = val
}

// }}}
// {{{ func(d *Data) refChange()

func (d *Data) refChange(val int) {
	d.val = val // alt: (*d).val = val
}

// }}}
// {{{ func(d *Data) String()

func (d *Data) String() string {
	d.prints += 1
	return fmt.Sprintf("Data{val = %d, prints = %d}", d.val, d.prints)
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	var x int = 0
	prefix := "     main():"
	fmt.Println(prefix, "x =", x, "/ &x =", &x)
	valChange(x)
	fmt.Println(prefix, "x =", x, "/ &x =", &x)
	refChange(&x)
	fmt.Println(prefix, "x =", x, "/ &x =", &x)
	fmt.Println("--------")
	dv := Data{1, 0}
	fmt.Printf("%s %s -- initial\n", prefix, &dv)
	dv.valChange(2)
	fmt.Printf("%s %s -- dv.valChange(2) applied\n", prefix, &dv)
	dv.refChange(3)
	fmt.Printf("%s %s -- dv.refChange(3) applied\n", prefix, &dv)
	fmt.Println("--------")
	dp := &Data{6, 0}
	fmt.Printf("%s %s -- initial\n", prefix, dp)
	dp.valChange(7)
	fmt.Printf("%s %s -- dp.valChange(7) applied\n", prefix, dp)
	dp.refChange(8)
	fmt.Printf("%s %s -- dp.refChange(8) applied\n", prefix, dp)
}

// }}}
