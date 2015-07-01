package hashers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHashers tests the hashers.
func TestMD5PasswordHasher(t *testing.T) {
	hashers := []PasswordHasher{
		&MD5PasswordHasher{},
		&SHA1PasswordHasher{},
		&PBKDF2SHA1PasswordHasher{},
		&PBKDF2PasswordHasher{},
	}

	for _, hasher := range hashers {
		password := "secret"
		encoded := hasher.Encode(password, "mysalt")
		assert.True(t, hasher.Verify(password, encoded))
	}
}
