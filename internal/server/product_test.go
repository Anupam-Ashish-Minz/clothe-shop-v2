package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullCheckProduct(t *testing.T) {
	assert := assert.New(t)
	p1, err := CheckProduct("mock product #1", "male", "this is a mock product", "5000")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("mock product #1", p1.Name)
	assert.Equal("male", p1.Gender)
	assert.Equal("this is a mock product", p1.Description)
	assert.Equal(500000, p1.Price)

	p2, err := CheckProduct("", "", "", "")
	assert.NotEqual(nil, err)
	assert.Equal("", p2.Name)
	assert.Equal("", p2.Gender)
	assert.Equal("", p2.Description)
	assert.Equal(0, p2.Price)

	p3, err := CheckProduct("", "male", "this is a mock product", "5000")
	assert.NotEqual(nil, err)
	assert.Equal("", p3.Name)
	assert.Equal("", p3.Gender)
	assert.Equal("", p3.Description)
	assert.Equal(0, p3.Price)

	p4, err := CheckProduct("mock product #1", "male", "", "5000")
	assert.Equal(nil, err)
	assert.Equal("mock product #1", p4.Name)
	assert.Equal("male", p4.Gender)
	assert.Equal("", p4.Description)
	assert.Equal(500000, p4.Price)

	p5, err := CheckProduct("mock product #1", "male", "this is a mock product", "4999.99")
	assert.Equal(nil, err)
	assert.Equal("mock product #1", p5.Name)
	assert.Equal("male", p5.Gender)
	assert.Equal("this is a mock product", p5.Description)
	assert.Equal(499999, p5.Price)

	p6, err := CheckProduct("mock product #1", "male", "this is a mock product", "$5000")
	assert.NotEqual(nil, err)
	assert.Equal("", p6.Name)
	assert.Equal("", p6.Gender)
	assert.Equal("", p6.Description)
	assert.Equal(0, p6.Price)
}
