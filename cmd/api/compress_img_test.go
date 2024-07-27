package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertImg(t *testing.T) {
	err := compressImgCli("img.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestChangeExtension(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(changeExtension("input.png"), "input.jpg")
	assert.Equal(changeExtension("input-name.png"), "input-name.jpg")
	assert.Equal(changeExtension("input.name.png"), "input.name.jpg")
	assert.Equal(changeExtension("input.name.another.png"), "input.name.another.jpg")
	assert.Equal(changeExtension(""), "")
	assert.Equal(changeExtension(".png"), "")
}
