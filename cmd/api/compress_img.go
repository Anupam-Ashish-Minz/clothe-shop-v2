package main

/*
#include <stdio.h>
#include <stdlib.h>

void print_message(const char *message) {
	printf("%s\n", message);
}
*/
import "C"

import (
	"unsafe"
)

func compressImg() error {
	message := C.CString("hello world")
	defer C.free(unsafe.Pointer(message))
	C.print_message(message)

	return nil
}
