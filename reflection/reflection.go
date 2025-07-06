//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Wed Oct 18 04:50:25 PM UTC 2023
//

// {{{ package

package reflection

// }}}
// {{{ imports

import (
	"fmt"
	"reflect"
)

// }}}
// ---- types ----
// {{{ type User struct

type User struct {
	Name string
	Age  int
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	var x float64 = 5.22
	var u User = User{"Albert", 55}
	fmt.Println("x =", x)
	fmt.Printf("u = %+v\n", u)
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Printf("v = %+v (%T) -> %s\n", v, v, v.String())
	fmt.Printf("t = %+v (%T) -> %s\n", t, t, t.String())
}

// }}}
