package main

// #cgo CFLAGS: -I ../../pkg/raylib/src -O2 -Wall -Wextra
// #cgo LDFLAGS: -static-libgcc -L ../../pkg/raylib/src -lraylib -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <raylib.h>
// static void MyPrint(char *s) { // Because go does not support variadic type
// 	printf("%s\n", s);
// }
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from stdio")
	C.MyPrint(cs)
	C.InitWindow(600, 600, C.CString("Hii"))
	C.free(unsafe.Pointer(cs))
}
