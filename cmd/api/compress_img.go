package main

/*
#cgo CFLAGS: -I/usr/include/ImageMagick-7 -fopenmp -DMAGICKCORE_HDRI_ENABLE=1 -DMAGICKCORE_QUANTUM_DEPTH=16 -DMAGICKCORE_CHANNEL_MASK_DEPTH=32
#cgo LDFLAGS: -lMagickWand-7.Q16HDRI -lMagickCore-7.Q16HDRI -lMagickWand-7.Q16HDRI -lMagickCore-7.Q16HDRI
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
