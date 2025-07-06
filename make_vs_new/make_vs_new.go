//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Tue Oct 10 05:20:23 PM UTC 2023
//

// {{{ package

package make_vs_new

// }}}
// {{{ imports

import "fmt"

// }}}
// ---- funcs ----
// {{{ func newInt1()

func newInt1() *int {
	return new(int)
}

// }}}
// {{{ func newInt1()

func newInt2() *int {
	var i int
	return &i
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	i1 := newInt1()
	i2 := newInt2()
	*i1 = 11
	*i2 = 12
	fmt.Printf("i1 = %d -> %T\n", *i1, i1) // *int
	fmt.Printf("i2 = %d -> %T\n", *i2, i2) // *int
	// channels
	c_make := make(chan int)
	c_new := new(chan int)
	fmt.Printf("c_make -> %T\n", c_make) // chan int
	fmt.Printf("c_new --> %T\n", c_new)  // *chan int
	// slices
	s_make := make([]int, 10)
	s_new := new([]int)
	fmt.Printf("s_make -> %T (len=%d, cap=%d)\n", s_make, len(s_make), cap(s_make)) // []int
	fmt.Printf("s_new --> %T (len=%d, cap=%d)\n", s_new, len(*s_new), cap(*s_new))  // *[]int
	// maps
	m_make := make(map[string]int, 20)
	m_make["a"] = 1
	m_new := new(map[string]int)
	// (*m_new)["b"] = 2  // panic: assignment to entry in nil map
	m_raw := map[string]int{"a": 1, "b": 2}
	fmt.Printf("m_make -> %T (len=%d)\n", m_make, len(m_make)) // map[string]int
	fmt.Printf("m_new --> %T (len=%d)\n", m_new, len(*m_new))  // *map[string]int
	fmt.Printf("m_raw --> %T (len=%d)\n", m_raw, len(m_raw))   // map[string]int
}

// }}}
