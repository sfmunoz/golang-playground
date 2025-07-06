//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/golang-playground
// Date: Wed Oct 18 03:55:37 PM UTC 2023
//

// {{{ package

package c_call

// }}}
// {{{ C code

// #cgo CFLAGS: -g -Wall -Wextra
// #include <stdlib.h>
// #include <stdio.h>
// int c_call()
// {
//   printf("running C code...\n");
//   return 0;
// }
// int my_age(const char *name, int age, char *out)
// {
//   int n;
//   n = sprintf(out, "Hello '%s', you are '%d' years old", name, age);
//   return n;
// }
import "C"

// }}}
// {{{ imports

import (
	"fmt"
	"unsafe"
)

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	// ----
	fmt.Println("running Go code...")
	ret := C.c_call()
	fmt.Println("ret =", ret)
	// ----
	name := C.CString("John")
	defer C.free(unsafe.Pointer(name))
	age := C.int(35)
	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))
	size := C.my_age(name, age, (*C.char)(ptr))
	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
	// ----
}

// }}}
