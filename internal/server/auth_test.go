package server

import "testing"

func TestJwtTokens(t *testing.T) {
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

	if parsedUserID != userID {
		t.Fatal("inputs and outputs don't match")
	}
}
