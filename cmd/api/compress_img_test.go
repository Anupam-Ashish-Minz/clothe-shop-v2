package main

import (
	"testing"
)

func TestCompressImg(t *testing.T) {
	err := compressImg()
	if err != nil {
		t.Fatal(err)
	}
}
