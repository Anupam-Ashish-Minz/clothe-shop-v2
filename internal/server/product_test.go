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
	assert.Equal(p1.Name, "mock product #1")
	assert.Equal(p1.Gender, "male")
	assert.Equal(p1.Description, "this is a mock product")
	assert.Equal(p1.Price, 500000)

	p2, err := CheckProduct("", "", "", "")
	assert.NotEqual(err, nil)
	assert.Equal(p2.Name, "")
	assert.Equal(p2.Gender, "")
	assert.Equal(p2.Description, "")
	assert.Equal(p2.Price, 0)

	p3, err := CheckProduct("", "male", "this is a mock product", "5000")
	assert.NotEqual(err, nil)
	assert.Equal(p3.Name, "")
	assert.Equal(p3.Gender, "")
	assert.Equal(p3.Description, "")
	assert.Equal(p3.Price, 0)

	p4, err := CheckProduct("mock product #1", "male", "", "5000")
	assert.Equal(err, nil)
	assert.Equal(p4.Name, "mock product #1")
	assert.Equal(p4.Gender, "male")
	assert.Equal(p4.Description, "")
	assert.Equal(p4.Price, 500000)

	p5, err := CheckProduct("mock product #1", "male", "this is a mock product", "4999.99")
	assert.Equal(err, nil)
	assert.Equal(p5.Name, "mock product #1")
	assert.Equal(p5.Gender, "male")
	assert.Equal(p5.Description, "this is a mock product")
	assert.Equal(p5.Price, 499999)

	p6, err := CheckProduct("mock product #1", "male", "this is a mock product", "$5000")
	assert.NotEqual(err, nil)
	assert.Equal(p6.Name, "")
	assert.Equal(p6.Gender, "")
	assert.Equal(p6.Description, "")
	assert.Equal(p6.Price, 0)
}
