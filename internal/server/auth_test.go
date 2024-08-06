package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	assert := assert.New(t)
	tokenString, err := createToken(int64(155), []byte("abcd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNTV9.nDzTEPrdzRwxhHnMnV0quPhKSjUqrRBwnxAeq7GN3Nw", tokenString)

	tokenString, err = createToken(int64(0), []byte("abcd"))
	if err == nil {
		t.Fatal("error should not be null")
	}
	assert.Equal("", tokenString)

	tokenString, err = createToken(int64(0), []byte(""))
	if err == nil {
		t.Fatal("error should not be null")
	}
	assert.Equal("", tokenString)
}

func TestParseToaken(t *testing.T) {
	assert := assert.New(t)
	secret := []byte("abcd")
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNTV9.nDzTEPrdzRwxhHnMnV0quPhKSjUqrRBwnxAeq7GN3Nw"
	parsedUserID, err := parseToken(tokenString, secret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(int64(155), parsedUserID)
}

func TestJwtTokens(t *testing.T) {
	assert := assert.New(t)
	userID := int64(155)
	secret := []byte("thisisthekey")
	tokenString, err := createToken(userID, secret)
	if err != nil {
		t.Fatal(err)
	}
	parsedUserID, err := parseToken(tokenString, secret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(userID, parsedUserID)
}
