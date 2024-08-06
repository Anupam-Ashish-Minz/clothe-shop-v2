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
	parsedUserID, err := parseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNTV9.nDzTEPrdzRwxhHnMnV0quPhKSjUqrRBwnxAeq7GN3Nw", []byte("abcd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(int64(155), parsedUserID)

	parsedUserID, err = parseToken("", []byte("abcd"))
	assert.Equal("token is malformed: token contains an invalid number of segments", err.Error())
	assert.Equal(int64(0), parsedUserID)

	parsedUserID, err = parseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", []byte("abcd"))
	assert.Equal("token signature is invalid: signature is invalid", err.Error())
	assert.Equal(int64(0), parsedUserID)

	parsedUserID, err = parseToken("eyJhbGciOiJQUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.Lfe_aCQme_gQpUk9-6l9qesu0QYZtfdzfy08w8uqqPH_gnw-IVyQwyGLBHPFBJHMbifdSMxPjJjkCD0laIclhnBhowILu6k66_5Y2z78GHg8YjKocAvB-wSUiBhuV6hXVxE5emSjhfVz2OwiCk2bfk2hziRpkdMvfcITkCx9dmxHU6qcEIsTTHuH020UcGayB1-IoimnjTdCsV1y4CMr_ECDjBrqMdnontkqKRIM1dtmgYFsJM6xm7ewi_ksG_qZHhaoBkxQ9wq9OVQRGiSZYowCp73d2BF3jYMhdmv2JiaUz5jRvv6lVU7Quq6ylVAlSPxeov9voYHO1mgZFCY1kQ", []byte("abcd"))
	assert.Equal("token is unverifiable: error while executing keyfunc: invalid signing method", err.Error())
	assert.Equal(int64(0), parsedUserID)

	parsedUserID, err = parseToken("eyJhbGciOiJIUzI1NInR5cCI6IkpXVCJ9.eyJVyX2lkIjoxNTV9.nDzTdzRwxhHnMnV0quPhKSjUqrRBwnxAeq7GN3Nw", []byte("abcd"))
	assert.Equal("token is malformed: could not base64 decode header: illegal base64 data at input byte 32", err.Error())
	assert.Equal(int64(0), parsedUserID)

	parsedUserID, err = parseToken("hello world", []byte("abcd"))
	assert.Equal("token is malformed: token contains an invalid number of segments", err.Error())
	assert.Equal(int64(0), parsedUserID)

	parsedUserID, err = parseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNTV9.nDzTEPrdzRwxhHnMnV0quPhKSjUqrRBwnxAeq7GN3Nw", []byte(""))
	assert.Equal("token signature is invalid: signature is invalid", err.Error())
	assert.Equal(int64(0), parsedUserID)

	parsedUserID, err = parseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTVkIn0.P3Cho5py4T9v-Imu1zisCzPFHvMBHZF_FjjskL2iIhg", []byte("abcd"))
	assert.Equal("failed to parse claims: float conversion", err.Error())
	assert.Equal(int64(0), parsedUserID)
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
