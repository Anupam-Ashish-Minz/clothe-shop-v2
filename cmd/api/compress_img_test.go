package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeExtension(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(changeExtension("input.png"), "input.jpg")
	assert.Equal(changeExtension("input-name.png"), "input-name.jpg")
	assert.Equal(changeExtension("input.name.png"), "input.name.jpg")
	assert.Equal(changeExtension("input.name.another.png"), "input.name.another.jpg")
	assert.Equal(changeExtension(""), "")
	assert.Equal(changeExtension(".png"), "")
	assert.Equal(changeExtension("/home/anupam/work/playground/web/htmx/clothe-shop-v2/tmp/data/output.png"),
		"/home/anupam/work/playground/web/htmx/clothe-shop-v2/tmp/data/output.jpg")
}

func TestConvertImg(t *testing.T) {
	err := compressImgCli("/home/anupam/work/playground/web/htmx/clothe-shop-v2/tmp/data/output.png")
	if err != nil {
		t.Fatal(err)
	}
}
